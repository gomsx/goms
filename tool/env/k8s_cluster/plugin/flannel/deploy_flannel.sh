#!/bin/bash
# set -xe

PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD

# docker pull quay.io/coreos/flannel:v0.10.0-amd64 

kubectl apply -f $PWD/kube-flannel.yml

