# CycloLab Microservice

## How to Run

1. Install Go (>=1.20)

2. Set the environment variable `APP_ENV` to choose the environment (demo, dev, prod). Default: `dev`

3. Configure environment-specific variables using a `.env` file inside the `configs` folder (optional), for example:

```env
SERVER_PORT=8080
LOG_LEVEL=info
```

4. Download dependencies:

```bash
go mod tidy
```

5. Run the microservice:

```bash
go run cmd/cyclolab/main.go
```
