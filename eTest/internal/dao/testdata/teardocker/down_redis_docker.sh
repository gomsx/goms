#!/bin/bash
set -x
set -e

set +x
echo "=================== down_redis_docker ========================"
set -x

## rm docker
# docker stop redistest
docker rm -f redistest

