#!/bin/bash
# set -xe

PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD

kubectl delete -f $PWD/mysql-deploy.yaml

kubectl delete -f $PWD/redis-sts.yaml

kubectl delete -f $PWD/user-deploy.yaml