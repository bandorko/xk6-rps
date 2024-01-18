## xk6-rps

This is an experimental extension for k6. Its purpose is to set the request limiter RPS to float value, because --rps flag accepts only integer value. https://k6.io/docs/using-k6/k6-options/reference/#rps

### build
```
xk6 build --with github.com/bandorko/xk6-rps
```

### usage

```
RPS=1.5 ./k6 run script.js
```