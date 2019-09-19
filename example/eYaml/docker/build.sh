#!/bin/bash
APP_IMAGE="eYaml"
DOCKER_IMAGE="eyaml"

go build -o $APP_IMAGE ../
ls -l ./$APP_IMAGE ../
chmod +x ./$APP_IMAGE 

docker build -t $DOCKER_IMAGE .
docker run -it $DOCKER_IMAGE

echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin 
docker tag $DOCKER_IMAGE $DOCKER_USERNAME/$DOCKER_IMAGE
docker push $DOCKER_USERNAME/$DOCKER_IMAGE
docker images