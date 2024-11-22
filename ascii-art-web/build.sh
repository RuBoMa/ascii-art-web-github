#!/bin/bash

# Variables
IMAGE_NAME="ascii-art-web"
TAG="latest"

# Build the docker image
echo "Building Docker image: $IMAGE_NAME:$TAG..."
docker build -t $IMAGE_NAME:$TAG .

# Start container on port 8080
echo "Starting container: $IMAGE_NAME:$TAG..."
docker run -p 8080:8080 --name asciiweb $IMAGE_NAME