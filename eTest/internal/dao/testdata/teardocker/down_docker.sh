#!/bin/bash
set -x
# set -e

# .sh所在目录
PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD

#
bash $PWD/down_mysql_docker.sh
bash $PWD/down_redis_docker.sh

