import http from 'k6/http';
import {check} from 'k6'

const BASE_URL = "http://localhost:8080"
//k6 run tests/k6_test/api_test_k6.js --vus 30 --duration 60s
export default function () {
  const res = http.get(BASE_URL+'/test');

  check(res, {
    'should return 200': (r) => r.status = 200
  })
}



