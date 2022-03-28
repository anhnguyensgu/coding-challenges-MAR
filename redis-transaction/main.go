package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-19548.c1.ap-southeast-1-1.ec2.cloud.redislabs.com:19548",
		Password: "6g62vFHHBuhi7fkyL8N2Q7ydHp2GJhz2", // no password set
	})

	tags := map[string]float64{
		"python":   3,
		"memcache": 1,
		"rust":     2,
		"c":        1,
		"redis":    1,
		"software": 1,
		"docker":   1,
		"go":       1,
		"linux":    1,
		"flask":    1,
	}

	for tag, score := range tags {
		// bs, _ := json.Marshal(tags)
		_, err := rdb.ZAdd("tags", redis.Z{-score, tag + time.Now().String()}).Result()
		if err != nil {
			log.Fatalf("Error adding %s", err)
		}
	}

	rdb.ZRemRangeByRank("tags", 10, -1)

	vals, err := rdb.ZRange("tags", 0, -1).Result()
	if err != nil {
		panic(err)
	}

	for _, val := range vals {
		fmt.Println(val)
	}
}
