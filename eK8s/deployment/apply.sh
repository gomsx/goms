#!/bin/bash
# set -xe

PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD

kubectl apply -f $PWD/mysql-sts.yaml --namespace="$1"

kubectl apply -f $PWD/redis-sts.yaml --namespace="$1"

kubectl apply -f $PWD/user-deploy.yaml --namespace="$1"

