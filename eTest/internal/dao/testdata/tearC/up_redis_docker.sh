#!/bin/bash
set -x
# set -e

set +x
echo "=================== up_redis_docker ========================"
set -x

## .sh所在目录
DIR=$(cd "$(dirname "$0")";pwd)
echo $DIR

#run docker
docker --version

docker pull redis
docker run --name redistest -p 26379:6379 -d redis

#ps docker
docker ps | grep redistest
