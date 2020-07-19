#!/bin/bash
set -x

# fluxctl
wget https://github.com/fluxcd/flux/releases/download/1.20.0/fluxctl_linux_amd64
mv fluxctl_linux_amd64 fluxctl 
chmod +x fluxctl 
cp fluxctl /usr/local/bin/

# minikube,安装socat
apt install -y socat

#

