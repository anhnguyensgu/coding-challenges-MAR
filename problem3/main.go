package main

import (
	"log"
	"problem3/web-service/mgconfig"
	"problem3/web-service/repository"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/qinains/fastergoding"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateEventRequest(event repository.Event) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(event)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func main() {
	fastergoding.Run()
	engine := html.New("./view", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/", "./view")

	client, ctx := mgconfig.InitializeMongoConnection()
	locationDatabase := client.Database("location-app")
	eventCollection := locationDatabase.Collection("event")
	visitorStatsCollection := locationDatabase.Collection("vistor-stats")
	evenRepository := repository.NewEventRepository(eventCollection)
	visitStatsRepository := repository.NewVistorStatsRepository(visitorStatsCollection)
	defer client.Disconnect(ctx)

	app.Get("/api/events", func(c *fiber.Ctx) error {
		events := evenRepository.GetAll(repository.PaginationParameter{Limit: 100, Skip: 0, SortField: "createdAt", Assending: -1})
		return c.JSON(events)
	})

	app.Post("/api/events", func(c *fiber.Ctx) error {
		event := &repository.Event{}
		if err := c.BodyParser(event); err != nil {
			log.Fatal(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		event.IpAddress = c.IP()

		errors := ValidateEventRequest(*event)
		if errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errors)

		}
		evenRepository.SaveEvent(*event)
		go visitStatsRepository.IncreaseVistorCount(*event)
		return c.JSON(event)
	})

	app.Get("/api/stats", func(c *fiber.Ctx) error {
		stats := visitStatsRepository.GetVisitStats()
		return c.JSON(stats)
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	log.Fatal(app.Listen(":3001"))
}
