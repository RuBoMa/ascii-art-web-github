#!/bin/bash

# Variables
IMAGE_NAME="ascii-art-web"
TAG="latest"

# Build the docker image
echo "Building Docker image: $IMAGE_NAME:$TAG..."
docker build -t $IMAGE_NAME:$TAG .