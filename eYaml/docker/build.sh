#!/bin/bash
APP_IMAGE="eYaml"
DOCKER_IMAGE="eyaml"

go build -o $APP_IMAGE ../cmd
ls -l ./$APP_IMAGE ../
chmod +x ./$APP_IMAGE 

# 下面 ../ 表示构建环境(目录)  
docker build -t $DOCKER_IMAGE -f dockerfile  ../    
docker run -it $DOCKER_IMAGE

echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin 
docker tag $DOCKER_IMAGE $DOCKER_USERNAME/$DOCKER_IMAGE
docker push $DOCKER_USERNAME/$DOCKER_IMAGE
docker images

#clear
rm $APP_IMAGE