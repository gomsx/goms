#!/bin/bash
# set -xe

# work dir
WD="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "--> work dir: ${WD}"

# work
kubectl apply -f ${WD}/2_3_1/recommended.yaml
kubectl apply -f ${WD}/2_3_1/dashboard-adminuser.yaml

kubectl patch svc -n kubernetes-dashboard kubernetes-dashboard -p '{"spec": {"type": "NodePort","ports": [{"port":443,"nodePort":30443}]}}'
