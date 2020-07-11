#!/bin/bash
# set -xe

PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD

kubectl apply -f $PWD/clientmock-deploy.yaml --namespace="$1"
kubectl apply -f $PWD/clientmock-svc.yaml --namespace="$1"

