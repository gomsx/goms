#!/bin/bash
set -x
set -e

set +x
echo "==> up docker"
set -x

pwdx=$(
	cd "$(dirname "$0")"
	pwd
)

## down
bash $pwdx/down_docker.sh

## run
docker run --name mysqltest -p 23306:3306 -d gomshub/mysqltest:v2.2.1
docker run --name redistest -p 26379:6379 -d gomshub/redistest:v1.4.2

## ps
docker ps | grep mysqltest
docker ps | grep redistest

## wait for docker init
sleep 35s

## log
docker logs mysqltest
docker logs redistest
