#!/bin/bash
set -x

# 查看
kubectl get all -n istio-system | grep grafana

# 集群内访问
# curl http://localhost:3000

# path 成 nodeport
# kubectl patch svc -n istio-system grafana -p '{"spec": {"type": "NodePort"}}'
kubectl patch svc -n istio-system grafana -p '{"spec": {"type": "NodePort","ports": [{"port":3000,"nodePort":31133}]}}'

# 查看
kubectl get all -n istio-system | grep grafana

# 集群外访问 url
# http://120.79.1.69:31192/grafana/