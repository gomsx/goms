#!/bin/bash
set -x
set -e

# work dir
WD="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "--> work dir: ${WD}"

# get config
CF="$WD/../config.yaml"
version="$(sed -n "s/\(version:\)\(.*\)/\2/p" ${CF} | tr -d ' ')"

# version
kubeadm_version="${version}-00"
kubelet_version="${version}-00"

# apt gpg
curl -s https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | apt-key add -

# apt config
cat <<EOF >/etc/apt/sources.list.d/kubernetes.list
deb https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main
EOF
sudo apt update

# apt install
sudo apt install kubeadm=${kubeadm_version} -y --allow-downgrades
sudo apt install kubelet=${kubelet_version} -y --allow-downgrades
