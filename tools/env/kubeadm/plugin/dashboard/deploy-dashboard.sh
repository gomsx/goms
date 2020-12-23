#!/bin/bash
# set -xe

pwdx=$(
	cd "$(dirname "$0")"
	pwd
)
echo $pwdx

kubectl apply -f $pwdx/2_0_0/recommended-my.yaml
kubectl apply -f $pwdx/2_0_0/dashboard-adminuser.yaml

