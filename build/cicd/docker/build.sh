#!/bin/bash
set -e
set -u
set -x

#define
DOCKER_IMAGE="ubuntu"

#build docker image
docker build -t $DOCKER_IMAGE .
docker images

#push docker image
echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin 
docker tag $DOCKER_IMAGE $DOCKER_USERNAME/$DOCKER_IMAGE
docker push $DOCKER_USERNAME/$DOCKER_IMAGE