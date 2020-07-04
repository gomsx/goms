#!/bin/bash
# set -xe

PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD

kubectl apply -f $PWD/mysql-svc.yaml

kubectl apply -f $PWD/redis-svc.yaml

kubectl apply -f $PWD/user-svc.yaml

