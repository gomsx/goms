#!/bin/bash
set -xe

# 1
kubectl create namespace goms

# mysql
# 1
kubectl apllay -f mysql-deploy.yaml --record
kubectl get rs,pod,deploy,svc -n goms
# 2
kubectl apllay -f mysql-svc.yaml --record
kubectl get rs,pod,deploy,svc -n goms
# 3
kubectl describe pod mysql-deploy -n goms
mysql -h 192.168.43.204 -P 31001 -u root -p

# redis
# 1
kubectl apllay -f redis-sts.yaml --record
kubectl get rs,pod,deploy,sts,svc -n goms
# 2
kubectl apllay -f redis-svc.yaml --record
kubectl get rs,pod,deploy,sts,svc -n goms
# 3
kubectl describe pod redis-deploy -n goms
redis-cli -h 192.168.43.204 -p 31002


# log
kubectl logs -n goms service/redis-svc


