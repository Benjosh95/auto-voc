#!/bin/bash

# Build and run the Docker image
docker-compose up --build -d

# Clean up dangling images
docker image prune -f

# Clean up unused containers
docker container prune -f
