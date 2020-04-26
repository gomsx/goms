#!/bin/bash
set -xe

sudo sysctl net.bridge.bridge-nf-call-iptables=1
wget https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml  2096  sed 's/quay.io\/coreos/registry.cn-beijing.aliyuncs.com\/imcto/g'
sudo kubectl apply -f kube-flannel.yml