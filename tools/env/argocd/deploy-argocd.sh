#!/bin/bash
set -x


kubectl create namespace argocd
kubectl apply -n argocd -f 2_0_0/install.yaml

# kubectl patch svc -n argocd argocd-server -p '{"spec": {"type": "NodePort","ports": [{"port":80,"nodePort":31140}]}}'
kubectl patch svc -n argocd argocd-server -p '{"spec": {"type": "NodePort","ports": [{"port":80,"nodePort":31140},{"port":443,"nodePort":31141}]}}'
