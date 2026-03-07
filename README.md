# microboiler_grpc_server

Boilerplate REST microservice template. Clean architecture, standard Go project layout.

## Structure

```
cmd/server/          — entrypoint
internal/
  config/            — env-based config
  domain/            — entities, interfaces (Repository)
  repository/        — infrastructure stub impl
  usecase/           — business logic
  grpc_handler/      — gRPC transport handlers
pkg/
  grpcserver/        — gRPC server + interceptors
api/                 — proto/contract references
```

## Run

```bash
cp .env.example .env
task run
```

## Build

```bash
task build
# binary: bin/server
```

## Test / Lint

```bash
task test
task lint
```

go run ./cmd/server
curl http://localhost:8080/api/users
curl -X POST http://localhost:8080/api/users -H 'Content-Type: application/json' -d '{"name":"Alice","email":"alice@example.com"}'
curl http://localhost:8080/api/users

## Health check

```bash

```
