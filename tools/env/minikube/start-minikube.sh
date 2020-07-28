#!/bin/bash

# clear
sudo rm -f /etc/kubernetes/admin.conf
sudo kubeadm reset
sudo rm -f $HOME/.kube $HOME/.minikube

# sudo minikube start --vm-driver=<driver_name>
minikube start --vm-driver=none --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers

# sudo minikube start/status/stop/delete
