#!/bin/bash
# set -xe

kubectl delete -f mysql-deploy.yaml
kubectl delete -f mysql-svc.yaml

kubectl delete -f redis-sts.yaml
kubectl delete -f redis-svc.yaml

kubectl delete -f user-deploy.yaml
kubectl delete -f user-svc.yaml