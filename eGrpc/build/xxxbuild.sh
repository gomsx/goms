#!/bin/bash
set -e
set -u
set -x

#define
APP_IMAGE="eGrpc"
DOCKER_IMAGE="egrpc"

#build app image
go build -o $APP_IMAGE ../cmd
ls -l ./$APP_IMAGE ../
chmod +x ./$APP_IMAGE 

#build docker image
docker build -t $DOCKER_IMAGE -f dockerfile  ../    
docker images

#push docker image
echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin 
docker tag $DOCKER_IMAGE $DOCKER_USERNAME/$DOCKER_IMAGE
docker push $DOCKER_USERNAME/$DOCKER_IMAGE

#clear 
rm $APP_IMAGE