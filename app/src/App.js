import { Paper } from '@mui/material';
import { useState } from 'react';
import "./App.css";
import { Chart } from './Chart';
import { getData } from './client';

const initialData = await getData()

function App() {
  const [data, setData] = useState(initialData)

  return (
    <div className="App">
      <header className="App-header">
        <div id="nav">
          <h1 id="homelink">codeCoach</h1>
        </div>
        <Paper
          elevation={3}
          sx={{ marginTop: 10, paddingTop: 10, marginX: 20 }}
        >
          <Chart data={data} />
        </Paper>
      </header>
    </div >
  );
}

export default App;
