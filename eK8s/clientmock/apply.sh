#!/bin/bash
# set -xe

PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD

kubectl apply -f $PWD/clientmock-deploy.yaml

kubectl apply -f $PWD/clientmock-svc.yaml