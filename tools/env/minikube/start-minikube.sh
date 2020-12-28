#!/bin/bash
set -x
set -e

# clear
sudo rm -f /etc/kubernetes/admin.conf
sudo rm -rf $HOME/.kube $HOME/.minikube

# local
# minikube start --vm-driver=none --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers

# wsl
if [ ! -d "/sys/fs/cgroup/systemd" ];then
    sudo mkdir /sys/fs/cgroup/systemd
    sudo mount -t cgroup -o none,name=systemd cgroup /sys/fs/cgroup/systemd
fi

# docker
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kicbase:v0.0.14
minikube start --vm-driver=docker --base-image="minikube/kicbase" --image-mirror-country='cn'