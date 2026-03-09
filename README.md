[//]: # (start name:common)
microboiler
===========

Project boilerplate

[//]: # (start name:project type:add)

rest_server
-----------

Boilerplate REST microservice template. Clean architecture, standard Go project layout.

### Structure

```
api/                 — proto/contract references
cmd/server/          — entrypoint
internal/
  config/            — env-based config
  domain/            — entities, interfaces (Repository)
  repository/        — infrastructure stub impl
  rest_handler/      — REST transport handlers
  rest_models/       — REST models
  usecase/           — business logic
pkg/
  httpserver/        — http server + interceptors
```

### Run

```bash
task run
```

### Build

```bash
task build
# binary: bin/server
```

### Test / Lint

```bash
task test
task lint
```

### Example requests

```bash
curl http://localhost:8080/api/users
curl -X POST http://localhost:8080/api/users -H 'Content-Type: application/json' -d '{"name":"Alice","email":"alice@example.com"}'
curl http://localhost:8080/api/users
```
