import { CartesianGrid, Legend, Line, LineChart, ResponsiveContainer, Tooltip, XAxis, YAxis } from 'recharts';

export function Chart({ data }) {

    return (
        <ResponsiveContainer width="90%" height={300}>
            <LineChart
                width={500}
                height={300}
                data={data}
                margin={{
                    top: 5,
                    right: 30,
                    left: 20,
                    bottom: 5,
                }}
            >
                <CartesianGrid strokeDasharray="3 3" />
                <XAxis dataKey="date" />
                <YAxis />
                <Tooltip />
                <Legend />
                <Line type="monotone" dataKey="linesSubtracted" stroke="#8884d8" activeDot={{ r: 8 }} />
                <Line type="monotone" dataKey="linesAdded" stroke="#82ca9d" />

            </LineChart>
        </ResponsiveContainer>
    )
}