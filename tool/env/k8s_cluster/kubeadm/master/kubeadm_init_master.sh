#!/bin/bash
set -xe

sudo kubeadm reset
sudo kubeadm init --kubernetes-version=v1.15.1 --apiserver-advertise-address=192.168.43.201 --pod-network-cidr=10.244.0.0/16      

# rm -rf $HOME/.kube/*
# mkdir -p $HOME/.kube
# sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
# sudo chown $(id -u):$(id -g) $HOME/.kube/config
