#!/bin/bash
set -x
# set -e

## .sh所在目录
DIR=$(cd "$(dirname "$0")";pwd)
echo $DIR

sh -c "$DIR/down_docker.sh"

sh -c "$DIR/up_mysql_docker.sh"

sh -c "$DIR/up_redis_docker.sh"
