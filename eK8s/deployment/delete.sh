#!/bin/bash
# set -xe

PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD

kubectl delete -f $PWD/mysql-sts.yam --namespace="$1"l

kubectl delete -f $PWD/redis-sts.yam --namespace="$1"l

kubectl delete -f $PWD/user-deploy.yam --namespace="$1"l

