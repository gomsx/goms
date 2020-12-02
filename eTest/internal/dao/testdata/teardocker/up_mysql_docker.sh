#!/bin/bash
set -x
# set -e

set +x
echo "==================== up_mysql_docker ======================"
set -x

#run docker
docker --version
docker run \
    --name mysqltest \
    -p 23306:3imageimage
    -d \
    mysqltest

#ps docker
docker logs mysqltest
docker ps | grep mysqltest

