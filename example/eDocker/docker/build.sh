#!/bin/bash
#echo $PATH;which go # 查看环境变量
go build ../
ls -l ./eDocker ../
chmod +x ./eDocker #重要
docker build -t sfw/edocker .
docker run -it sfw/edocker

echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin 
docker tag sfw/edocker dockerxpub/edocker
docker push dockerxpub/edocker
docker images