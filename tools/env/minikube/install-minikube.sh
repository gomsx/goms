#!/bin/bash
set -x
set -e

# minikube
curl -Lo minikube https://github.com/kubernetes/minikube/releases/download/v1.12.0/minikube-linux-amd64

chmod +x minikube
sudo mkdir -p /usr/local/bin/
sudo install minikube /usr/local/bin/
