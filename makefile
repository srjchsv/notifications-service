# Load environment variables from .env file
include dev.env

export
.ONESHELL:

run:
	@export BROKER_HOST=localhost
	@go run cmd/myapp/main.go

container:
	docker compose up