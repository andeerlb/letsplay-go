FROM golang:1.25-rc-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG APP_ENV
ENV APP_ENV=${APP_ENV}

RUN go build -ldflags="-s -w" -o app cmd/letsplay/main.go

FROM debian:bullseye-slim

ARG APP_ENV
ENV APP_ENV=${APP_ENV}

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/app .
COPY configs/${APP_ENV}.yaml ./config.yaml

ENV APP_ENV=${APP_ENV}

EXPOSE 9090

CMD ["./app"]
