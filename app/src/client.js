import axios from 'axios'

export const client = axios.create({
    baseURL: 'http://localhost:8000'
})

export async function getData(window) {
    var url = '/data'
    if (window) {
        url = url + `?window=${window}`
    }

    const resp = await client.get(url)
    console.log(resp.data)
    return resp.data
}
