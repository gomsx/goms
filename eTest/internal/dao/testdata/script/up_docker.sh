#!/bin/bash
set -x
set -e

set +x
echo "==> up docker"
set -x

## test env
echo "--> TEST_WITH_BASE_SERVER: ${TEST_WITH_BASE_SERVER}"
tenv=${TEST_WITH_BASE_SERVER,,} #转小写
if [[ ${tenv} == "true" || ${tenv} == "yes" || ${tenv} == "ok" ]]; then
	echo "--> test env with base service"
	rm -rf ../configs
	cp -rf ../configs_default ../configs
	exit 0
fi
rm -rf ../configs
cp -rf ../configs_docker ../configs
echo "--> test env without base service, start container service now"

## down
bash ./down_docker.sh

## run
docker run --name mysqltest -p 23306:3306 -d dockerxpub/mysqltest:v2.2.4
docker run --name redistest -p 26379:6379 -d dockerxpub/redistest:v1.4.4

## ps
docker ps | grep mysqltest
docker ps | grep redistest

## wait for docker init
sleep 35s

## log
docker logs mysqltest
docker logs redistest
