#!/bin/bash

# Variables
IMAGE_NAME="ascii-art-web"
TAG="latest"

# Build the docker image
echo "Building Docker image: $IMAGE_NAME:$TAG..."
docker build --tag $IMAGE_NAME:$TAG .

# Start container on port 8080
echo "Starting container in Docker: $IMAGE_NAME:$TAG..."
docker run --detach --port 8080:8080 --name asciiweb $IMAGE_NAME