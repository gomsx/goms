#!/bin/bash
set -x

NS="$1"
[[ -z "$NS" ]] || NS="$(kubens -c)"

kubectl patch svc -n "$NS" user-svc -p '{"spec": {"type": "NodePort","ports": [{"port":8080,"nodePort":32023}]}}'
kubectl patch svc -n "$NS" redis-svc -p '{"spec": {"type": "NodePort","ports": [{"port":6379,"nodePort":32022}]}}'
kubectl patch svc -n "$NS" mysql-svc -p '{"spec": {"type": "NodePort","ports": [{"port":3306,"nodePort":32021}]}}'
