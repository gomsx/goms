#!/bin/bash
# set -xe

PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD

kubectl create configmap cm-user --from-file=$PWD/configs -n goms-ek8s
# kubectl describe configmaps cm-user -n goms-ek8s
# kubectl get configmaps cm-user -o yaml