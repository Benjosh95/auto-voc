# Log in to Docker Hub
docker login

# Build the Docker image
docker build -t auto-voc:1.0.0 .

# Tag the Docker image
docker tag auto-voc:1.0.0 <your-dockerhub-username>/auto-voc:1.0.0

# Push the Docker image to Docker Hub
docker push <your-dockerhub-username>/auto-voc:1.0.0
