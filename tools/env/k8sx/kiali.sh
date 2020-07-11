#!/bin/bash
set -x

# 查看
kubectl get all -n istio-system | grep kiali

# 集群内访问
# curl http://localhost:20001

# path 成 nodeport
# kubectl patch svc -n istio-system kiali -p '{"spec": {"type": "NodePort"}}'
kubectl patch svc -n istio-system kiali -p '{"spec": {"type": "NodePort","ports": [{"port":20001,"nodePort":31131}]}}'

# 查看
kubectl get all -n istio-system | grep kiali

# 集群外访问 url
# http://120.79.1.69:31192/kiali/
