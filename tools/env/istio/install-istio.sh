#!/bin/bash
set -x

#
sudo tar -C /usr/local -xvzf istio-1.6.4-linux-amd64.tar.gz 
echo "PATH=/usr/local/istio-1.6.4/bin:\$PATH" >> $HOME/.bashrc
source $HOME/.bashrc

#
istioctl profile list
istioctl manifest apply --set profile=demo

