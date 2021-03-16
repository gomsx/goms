#!/bin/bash
set -x

# 查看
kubectl get all -n istio-system | grep prometheus

# 集群内访问
# curl http://localhost:9090

# path 成 nodeport
# kubectl patch svc -n istio-system prometheus -p '{"spec": {"type": "NodePort"}}'
kubectl patch svc -n istio-system prometheus -p '{"spec": {"type": "NodePort","ports": [{"port":9090,"nodePort":31130}]}}'

# 查看
kubectl get all -n istio-system | grep prometheus

# 集群外访问 url
# http://120.79.1.69:31130/graph

