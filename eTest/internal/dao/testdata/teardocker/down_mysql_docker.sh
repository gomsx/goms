#!/bin/bash
set -x
set -e

set +x
echo "==================== down_mysql_docker ======================"
set -x

## rm docker
# docker stop mysqltest
docker rm -f mysqltest

