#!/bin/bash
# set -xe

pwdx=$(
	cd "$(dirname "$0")"
	pwd
)
echo $pwdx

kubectl apply -f $pwdx/1_2_0/kube-flannel.yaml

