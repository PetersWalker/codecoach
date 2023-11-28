import axios from 'axios'

export const client = axios.create({
    baseURL: 'http://localhost:8000'
})

export async function getData(data) {
    const resp = await client.get('/data', data)
    console.log(resp.data)
    return resp.data
}
