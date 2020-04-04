#!/bin/bash
set -x
# set -e

## .sh所在目录
PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD

# bash $PWD/down_docker.sh >&1 | tee $PWD/output.log
bash $PWD/up_mysql_docker.sh >&1 | tee -a $PWD/output.log
bash $PWD/up_redis_docker.sh >&1 | tee -a $PWD/output.log

set +x;echo " --------------- docker running ----------------";set -x

docker ps | grep mysqltest
docker ps | grep redistest

echo -e "\n\n"
