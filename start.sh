#!/bin/sh

ENV=""

for arg in "$@"
do
  case $arg in
    --env=*)
      ENV="${arg#*=}"
      shift
      ;;
  esac
done

if [ -z "$ENV" ]; then
  if ! command -v gum > /dev/null; then
    echo "Error: gum is not installed. Install it from https://github.com/charmbracelet/gum"
    exit 1
  fi

  ENV=$(gum choose "dev" "qa" "prod")
fi

case "$ENV" in
  dev|qa|prod)
    ;;
  *)
    echo "Invalid environment: $ENV"
    echo "Valid options: dev, qa, prod"
    exit 1
    ;;
esac

echo "Starting with APP_ENV=$ENV"

APP_ENV="$ENV" go run cmd/letsplay/main.go
