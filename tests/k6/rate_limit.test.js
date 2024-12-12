import http from 'k6/http'
import { sleep, check } from 'k6'

const AUTH_TOKEN = 'AUTH_TOKEN'
const API_KEY = 'API_KEY'

export let options = {
    vus: 1,
    duration: '6s',
    rps: 10
}

export default function() {
    const url = 'http://localhost:5700/api/v1/orders'
    const payload = JSON.stringify({
        menu_id: 'MDE5M2JhZWEtNDYzNi0xOTQ5LTNmYTUtNmIxNmQ1NTBlM2I0',
        payment_method: 'COD'
    })
    const params = {
        headers: {
            'Authorization': `Bearer ${AUTH_TOKEN}`,
            'x-api-key': `Key ${API_KEY}`
        }
    }

    const res = http.post(url, payload, params)
    
    check(res, {
        'is status 200, ok!': (res) => res.status === 200,
        'is status 429, limit reached!': (res) => res.status === 429
    })
}