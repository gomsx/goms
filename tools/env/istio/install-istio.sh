#!/bin/bash
set -x

version="1.9.1"

# get
wget https://github.com/istio/istio/releases/download/${version}/istio-${version}-linux-amd64.tar.gz

# install
sudo tar -C /usr/local -xvzf istio-${version}-linux-amd64.tar.gz
echo "export PATH=\"/usr/local/istio-${version}/bin:\$PATH\"" >> $HOME/.bashrc
source $HOME/.bashrc

# deploy
istioctl profile list
istioctl manifest apply --set profile=demo
