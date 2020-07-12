#!/bin/bash
# set -xe

PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD

kubectl create configmap cm-user --from-file=$PWD/configs -n "$1"
# kubectl describe configmaps cm-user -n goms
# kubectl get configmaps cm-user -o yaml

