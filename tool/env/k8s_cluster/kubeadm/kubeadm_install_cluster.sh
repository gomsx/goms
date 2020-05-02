#!/bin/bash
set -xe

curl -s https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | apt-key add -

cat <<EOF > /etc/apt/sources.list.d/kubernetes.list
deb https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main
EOF

sudo apt-get update
# sudo apt-get install -y kubeadm
sudo apt install kubelet=1.15.1-00
sudo apt install kubeadm=1.15.1-00
sudo apt install kubectl=1.15.1-00
