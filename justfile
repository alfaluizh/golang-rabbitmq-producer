# Starts the RabbitMQ server
infra:
  docker compose up -d

# Build and runs the api
run:
  go run cmd/api/main.go
