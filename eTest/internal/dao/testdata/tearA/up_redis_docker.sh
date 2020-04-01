#!/bin/bash
set -x
# set -e

## .sh所在目录
DIR=$(cd "$(dirname "$0")";pwd)
echo $DIR

#run docker
docker --version

docker pull redis
docker run --name redistest -p 31007:6379 -d redis

#delay
sleep 5 
docker ps | grep redistest

#ps docker
docker ps | grep redistest