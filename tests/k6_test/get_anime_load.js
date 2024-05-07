
import { check } from 'k6';
import http from 'k6/http';

const BASE_URL = "http://localhost:8080"

export const options = {
    stages:[
        {duration:'5s',target:10},
        {duration:'5s',target:10},
        {duration:'15s',target:100},
        {duration:'10s',target:100},
        {duration:'10s',target:0},
    ],
    thresholds:{
        http_req_duration:['p(95)<1000'],//95% das req devem responder em até 2s
        http_req_failed:['rate<0.09']//1% das req podem falhar
    }
}

export default function () {
    const res = http.get(BASE_URL + '/anime/1');

    console.log(res.body)
    check(res, {
        'should return 200': (r) => r.status === 200
    })
}