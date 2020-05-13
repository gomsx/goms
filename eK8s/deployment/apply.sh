#!/bin/bash
# set -xe

PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD

kubectl apply -f $PWD/mysql-deploy.yaml

kubectl apply -f $PWD/redis-sts.yaml

kubectl apply -f $PWD/user-deploy.yaml