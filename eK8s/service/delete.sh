#!/bin/bash
# set -xe

PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD

kubectl delete -f $PWD/mysql-svc.yaml

kubectl delete -f $PWD/redis-svc.yaml

kubectl delete -f $PWD/user-svc.yaml

