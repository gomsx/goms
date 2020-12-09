#!/bin/bash
# set -x
# set -e

# version
KUBEADM_VERSION=1.18.1-00
KUBELET_VERSION=1.18.1-00
KUBECTL_VERSION=1.18.1-00

# apt gpg
curl -s https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | apt-key add -

# apt config
cat <<EOF > /etc/apt/sources.list.d/kubernetes.list
deb https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main
EOF
sudo apt update

# apt install
sudo apt remove kubeadm kubelet kubectl -y
sudo apt install kubeadm=$KUBEADM_VERSION -y --allow-downgrades
sudo apt install kubelet=$KUBELET_VERSION -y --allow-downgrades
sudo apt install kubectl=$KUBECTL_VERSION -y --allow-downgrades

