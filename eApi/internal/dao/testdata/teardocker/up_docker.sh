#!/bin/bash
set -x
set -e

## .sh所在目录
PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD

## bash
bash $PWD/down_docker.sh >&1 | tee $PWD/output.log
bash $PWD/up_mysql_docker.sh >&1 | tee -a $PWD/output.log
bash $PWD/up_redis_docker.sh >&1 | tee -a $PWD/output.log

## echo
set +x
echo "==================== docker running ======================"
set -x

