#!/bin/bash
set -x
# set -e

## .sh所在目录
DIR=$(cd "$(dirname "$0")";pwd)
echo $DIR

bash $DIR/down_docker.sh >&1 | tee $DIR/output.log
bash $DIR/up_mysql_docker.sh >&1 | tee -a $DIR/output.log
bash $DIR/up_redis_docker.sh >&1 | tee -a $DIR/output.log

set +x;echo " --------------- docker running ----------------";set -x

docker ps | grep mysqltest
docker ps | grep redistest

echo -e "\n\n"
