.PHONY: build-local run-local stop-local logs-local \
	build-prod run-prod stop-prod logs-prod test lint run-host

LOCAL_COMPOSE = docker compose -f docker-compose.yml -f docker-compose.local.yml

# Local stack: source is bind-mounted and Air rebuilds on save, so there is no
# image rebuild. Only run `make build-local` when go.mod changes.
build-local:
	$(LOCAL_COMPOSE) build

run-local:
	$(LOCAL_COMPOSE) up

stop-local:
	$(LOCAL_COMPOSE) down

logs-local:
	$(LOCAL_COMPOSE) logs -f api

# Prod-style stack: the code is compiled into the image, so run
# `make build-prod` after code changes. `run-prod` starts it in the background.
build-prod:
	docker compose build

run-prod:
	docker compose up -d

stop-prod:
	docker compose down

logs-prod:
	docker compose logs -f api

test:
	go test -race ./...

lint:
	go vet ./...
	@test -z "$$(gofmt -l .)" || (echo "gofmt needed:"; gofmt -l .; exit 1)

# Runs the API on the host with no Docker; needs the env vars from .env exported
# (wired up in step 2).
run-host:
	go run ./cmd/api
