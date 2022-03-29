import { useEffect, useState } from 'react';
import logo from './logo.svg';
import './App.css';
import axios from 'axios';
import 'antd/dist/antd.min.css';

function App() {
  const [latestVistors, setLatestVisitor] = useState([]);
  const [mostVisitUser, setMostVisitUser] = useState([]);

  useEffect(() => {
    axios.get('/api/events').then((res) => setLatestVisitor(res.data));
    axios.get('/api/stats').then((res) => setMostVisitUser(res.data));
  }, []);

  return (
    <div>
      {/* <div>
      {latestVistors.map(({ IpAddress }) => (
        <p>{IpAddress}</p>
      ))}
      </div> */}
      <div>
        {mostVisitUser.map(({ IpAddress,Visit }) => (
          <p>{IpAddress}-{Visit}</p>
        ))}
      </div>
    </div>
  );
}

export default App;
