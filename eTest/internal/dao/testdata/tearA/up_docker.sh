#!/bin/bash
set -x
# set -e

## .sh所在目录
DIR=$(cd "$(dirname "$0")";pwd)
echo $DIR

source $DIR/up_mysql_docker.sh
source $DIR/up_redis_docker.sh