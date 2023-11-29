import { Box, FormControl, InputLabel, MenuItem, Paper, Select } from '@mui/material';
import { useState } from 'react';
import "./App.css";
import { Chart } from './Chart';
import { getData } from './client';


const initialData = await getData()

function App() {
  const [data, setData] = useState(initialData)
  const [dataWindow, setDataWindow] = useState('week')

  async function getFreshData(value) {
    const data = await getData(value)
    setData(data)
  }

  return (
    <div className="App">
      <header className="App-header">
        <div id="nav">
          <h1 id="homelink">codeCoach</h1>
        </div>
        <Paper
          elevation={3}
          sx={{ marginTop: 8, paddingTop: 8, marginX: 16 }}
        >

          <Box sx={{ minWidth: 120, marginBottom: 8, marginLeft: 8 }}>
            <FormControl sx={{ minWidth: 120 }}>
              <InputLabel id="demo-simple-select-label">Age</InputLabel>
              <Select
                labelId="demo-simple-select-label"
                id="demo-simple-select"
                value={dataWindow}
                label="Data window"
                onChange={(e) => {
                  getFreshData(e.target.value)
                  setDataWindow(e.target.value)
                }}
              >
                <MenuItem value={"week"}>week</MenuItem>
                <MenuItem value={"month"}>month</MenuItem>
                <MenuItem value={"year"}>year</MenuItem>
              </Select>
            </FormControl>
          </Box>

          <Chart data={data} />
        </Paper>
      </header>
    </div >
  );
}

export default App;
