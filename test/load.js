import http from 'k6/http';
import { check } from 'k6';

export const options = {
  stages: [
    { duration: '10s', target: 10 },
  ],
};

export default function () {
  const url = 'http://localhost:8080/logs';

  const payload = JSON.stringify({
    ts: Date.now(),
    level: 'info',
    message: `
LoremLorem Ipsum is simply dummy text of the printing and typesetting industry. 
Lorem Ipsum has been the industry standard dummy text ever since the 1500s, when
an unknown printer took a galley of type and scrambled it to make a type specimen book.
`,
    service: 'k6'
  });

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  for (let i = 0; i < 10000; i++) {
    const res = http.post(url, payload, params);

    check(res, {
      'status 201': (r) => r.status == 201,
    });
  }
}

