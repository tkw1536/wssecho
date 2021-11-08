# wssecho


![CI Status](https://github.com/tkw1536/wssecho/workflows/CI/badge.svg)

A quick dummy websocket implementation that echoes and can be used for testing.

Run it using:

```bash
go run . -bind localhost:8080
```

The client code is in [index.js](./index.js), the main server code in [main.go](./main.go).
