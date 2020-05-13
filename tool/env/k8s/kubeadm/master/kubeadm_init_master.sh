#!/bin/bash
set -xe

sudo kubeadm reset
sudo kubeadm init --kubernetes-version=v1.15.1 --apiserver-advertise-address=192.168.1.131 --pod-network-cidr=10.244.0.0/16      
