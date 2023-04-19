#!/bin/bash

if [[ $1 == "dev" ]]; then
     # Load environment variables from dev.env file
    source dev.env

    # Export environment variables
    export $(sed 's/=.*//' dev.env)

    go run cmd/myapp/main.go

elif [[ $1 == "compose" ]]; then
    # Load environment variables from .env file
    source .env

    # Export environment variables
    export $(sed 's/=.*//' .env)

    # Run docker compose
    docker compose up

else
    echo "Invalid argument. Please specify either 'dev' or 'compose'"
fi

# Stop the containers
docker compose down
