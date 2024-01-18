## xk6-rps

This is an PoC extension for k6. Its purpose is to set the request limiter RPS to float value, because --rps flag accepts only integer value. https://k6.io/docs/using-k6/k6-options/reference/#rps

### build
```
xk6 build --with github.com/bandorko/xk6-rps
```

### usage

### script.js :
```
import http from 'k6/http';
import 'k6/x/rps';

export let options = {
  vus: 10,
  duration: '1m',
};

export default function () {
  // GET #1
  http.get('https://test.k6.io/');
  // GET #2
  http.get('https://test.k6.io/');
  // GET #3
  http.get('https://test.k6.io/');
}
```

You can set the desired rps with the RPS environment variable
```
RPS=1.5 ./k6 run script.js
```