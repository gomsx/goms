#!/bin/bash
set -xe

sudo kubeadm init --apiserver-advertise-address=192.168.43.201 --pod-network-cidr=10.244.0.0/16      
