#!/bin/bash

imageName="notifications-service"

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

elif [[ $1 == "build" ]]; then
    # Build image
    docker build -t $imageName .

elif [[ $1 == "push" ]]; then
    # Tag image
    docker tag $imageName srjchsv/$imageName

    # Push image to remote registry
    docker push srjchsv/$imageName

else
    echo "Invalid argument"
fi
