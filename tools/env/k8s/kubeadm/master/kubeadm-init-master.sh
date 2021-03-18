#!/bin/bash
# set -e
set -x

# work dir
WD="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "--> work dir: ${WD}"

# get config
CF="$WD/../config.yaml"
version="$(sed -n "s/\(version:\)\(.*\)/\2/p" ${CF} | tr -d ' ')"
master_ip="$(sed -n "s/\(master_ip:\)\(.*\)/\2/p" ${CF} | tr -d ' ')"

# init
image_repo="registry.aliyuncs.com/google_containers"
sudo kubeadm reset
sudo kubeadm init --kubernetes-version=${version} --apiserver-advertise-address=${master_ip} --image-repository=${image_repo} --pod-network-cidr=10.244.0.0/16
