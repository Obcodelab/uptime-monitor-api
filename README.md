# Uptime Monitor API

A REST API in Go (Chi, PostgreSQL, Docker) that checks the uptime of user-supplied
URLs in the background and exposes monitors, check history and statistics.

> Work in progress: only the scaffold (`GET /healthz`) exists so far.

## Quick start

```sh
cp .env.example .env
docker compose up --build
curl localhost:8000/healthz
```

## Development

`make run-local` runs the stack with the source mounted and
[Air](https://github.com/air-verse/air) rebuilding on every save, so no image rebuild is
needed. Run `make build-local` only after changing `go.mod`. `make build-prod` and
`make run-prod` use the production-style image instead.

## Tests

```sh
make test
```
