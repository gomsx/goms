#!/bin/bash
set -xe

kubectl create configmap ek8sv3 --from-file=./configs -n ek8sv3
kubectl describe configmaps ek8sv3 -n ek8sv3
kubectl get configmaps ek8sv3 -o yaml