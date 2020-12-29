#!/bin/bash
set -x
set -e

set +x
echo "=============== down docker ================"
set -x

docker rm -f mysqltest
docker rm -f redistest
