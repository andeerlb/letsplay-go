# letsplay Microservice

Go microservice for creating players, which calls another microservice for persistence.

## Project Structure

* `cmd/letsplay/main.go`: Entry point
* `internal/config`: Configuration using Viper (supports demo/dev/prod profiles)
* `internal/logger`: Structured logger with Zap
* `internal/model`: Models and structs
* `internal/client`: HTTP client for the other microservice
* `internal/service`: Business logic
* `internal/handler`: HTTP handlers (controllers)
* `internal/router`: Route configuration
* `internal/server`: HTTP server setup
* `tests`: Unit and integration tests (not implemented here)

## How to Run

1. Install Go (>=1.20)

2. Set the environment variable `APP_ENV` to choose the environment (demo, dev, prod). Default: `dev`

3. Configure environment-specific variables using a `.env` file inside the `configs` folder (optional), for example:

```env
SERVER_PORT=8080
LOG_LEVEL=info
PLAYER_SERVICE_URL=http://localhost:9000
```

4. Download dependencies:

```bash
go mod tidy
```

5. Run the microservice:

```bash
go run cmd/letsplay/main.go
```

6. Test the endpoint:

```bash
curl -X POST http://localhost:8080/newplayer \
 -H 'Content-Type: application/json' \
 -d '{
     "email": "player@example.com",
     "password": "secret123",
     "primaryGame": {
         "type": "futebol",
         "position": "atacante"
     },
     "otherGame": [
         {"type": "basquete", "position": "armador"},
         {"type": "volei", "position": "levantador"}
     ]
 }'
```

## How to Expand

* Add unit and integration tests
* Authentication middleware (JWT/OAuth)
* Monitoring and metrics (Prometheus, OpenTelemetry)
* Dynamic configuration via Consul/Etcd
* Dockerfile and CI/CD pipeline
* Improved error handling and structured logging
