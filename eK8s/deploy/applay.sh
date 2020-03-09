#!/bin/bash
set -xe

kubectl apply -f mysql-deploy.yaml
kubectl apply -f mysql-svc.yaml

kubectl apply -f redis-sts.yaml
kubectl apply -f redis-svc.yaml

kubectl apply -f user-deploy.yaml
kubectl apply -f user-svc.yaml