#!/bin/bash
# set -xe

PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD

kubectl apply -f $PWD/2_0_0/recommended-my.yaml
kubectl apply -f $PWD/2_0_0/dashboard-adminuser.yaml

