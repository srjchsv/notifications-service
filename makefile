# Load environment variables from .env file
include .env

export

run:
	go run cmd/myapp/main.go