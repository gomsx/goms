#!/bin/bash
# set -xe

sudo rm -f /etc/kubernetes/admin.conf
sudo kubeadm reset
sudo kubeadm init --kubernetes-version=v1.18.1 --apiserver-advertise-address=172.18.194.179 --pod-network-cidr=10.244.0.0/16      
