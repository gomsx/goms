#!/bin/bash

set -x

# docker

# podman
# sudo apt -y  install software-properties-common
# sudo add-apt-repository -y ppa:projectatomic/ppa
# sudo apt -y install podman
# sudo -k -n podman version

# kubectl
# sudo apt install snap
# sudo snap install kubectl --classic
# kubectl version --client

# minikube
curl -Lo minikube https://github.com/kubernetes/minikube/releases/download/v1.12.0/minikube-linux-amd64

chmod +x minikube
sudo mkdir -p /usr/local/bin/
sudo install minikube /usr/local/bin/
