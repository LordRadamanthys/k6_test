
import { check } from 'k6';
import http from 'k6/http';
import { htmlReport } from "https://raw.githubusercontent.com/benc-uk/k6-reporter/main/dist/bundle.js";
const BASE_URL = "http://localhost:8080"

export const options = {
    vus: 100,
    duration: '10s',
    thresholds:{
        http_req_duration:['p(95)<900'],//95% das req devem responder em até 2s
        http_req_failed:['rate<0.01']//1% das req podem falhar
    }
}
export function handleSummary(data) {
    return {
      "summary.html": htmlReport(data),
    };
  }

export default function () {
    const res = http.get(BASE_URL + '/anime/1');

    console.log(res.body)
    check(res, {
        'should return 200': (r) => r.status === 200
    })
}