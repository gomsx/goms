#!/bin/bash
#echo $PATH;which go # 查看环境变量
go build ../
ls -l ./eYaml ../
chmod +x ./eYaml #重要
docker build -t sfw/eyaml .
docker run -it sfw/eyaml

echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin 
docker tag sfw/eyaml dockerxpub/eyaml
docker push dockerxpub/eyaml
docker images