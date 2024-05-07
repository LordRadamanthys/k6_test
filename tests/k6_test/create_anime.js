
import { check } from 'k6';
import http from 'k6/http';

const BASE_URL = "http://localhost:8080"

export default function () {
    const payload = JSON.stringify( {
		title: "Cowboy Bebop",
		title_japanese: "カウボーイビバップ",
		duration: "24 min per ep"
	})
    const params = {
        headers: {
          'Content-Type': 'application/json',
        }, 
      };
    const res = http.post(BASE_URL + '/anime',payload, params);

    check(res, {
        'should return 200': (r) => r.status === 200
    })
}