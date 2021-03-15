#!/bin/bash
# set -e
set -x

version="$(sed -n "s/\(version:\)\(.*\)/\2/p" ${WD}/config.yaml | tr -d ' ')"
master_ip="$(sed -n "s/\(master_ip:\)\(.*\)/\2/p" ${WD}/config.yaml | tr -d ' ')"

sudo kubeadm reset
sudo kubeadm init --kubernetes-version=${version} --apiserver-advertise-address=${master_ip} --pod-network-cidr=10.244.0.0/16
