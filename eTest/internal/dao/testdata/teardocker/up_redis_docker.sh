#!/bin/bash
set -x
# set -e

set +x
echo "=================== up_redis_docker ========================"
set -x

#run docker
docker run \
    --name redistest \
    -p 26379:6379 \
    -d \
    redistest
#ps docker
docker ps | grep redistest

