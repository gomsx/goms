#!/bin/bash
set -x

PF="test-"
NS=test-goms

kubectl patch svc -n "$NS" "$PF"user-svc -p '{"spec": {"type": "NodePort","ports": [{"port":8080,"nodePort":32023}]}}'
kubectl patch svc -n "$NS" "$PF"redis-svc -p '{"spec": {"type": "NodePort","ports": [{"port":6379,"nodePort":32022}]}}'
kubectl patch svc -n "$NS" "$PF"mysql-svc -p '{"spec": {"type": "NodePort","ports": [{"port":3306,"nodePort":32021}]}}'
