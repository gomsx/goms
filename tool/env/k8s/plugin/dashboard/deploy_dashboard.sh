#!/bin/bash
# set -xe

PWD=$(cd "$(dirname "$0")";pwd)
echo $PW

kubectl apply -f $PWD/2_0_0/recommended-my.yaml

# sudo kubectl get rs,pod,deploy,svc,ep -n kubernetes-dashboard

# sudo kubectl delete -f 2_0_0/recommended-my.yaml