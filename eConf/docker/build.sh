#!/bin/bash
APP_IMAGE="eConf"
DOCKER_IMAGE="econf"

go build -o $APP_IMAGE ../cmd
ls -l ./$APP_IMAGE ../
chmod +x ./$APP_IMAGE 

docker build -t $DOCKER_IMAGE -f dockerfile  ../    

echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin 
docker tag $DOCKER_IMAGE $DOCKER_USERNAME/$DOCKER_IMAGE
docker push $DOCKER_USERNAME/$DOCKER_IMAGE
docker images