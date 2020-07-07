#!/bin/bash

set -x

sudo tar -C /usr/local -xvzf istio-1.6.4-linux-amd64.tar.gz 
echo "PATH=/usr/local/istio-1.6.4:\$PATH" >> .bashrc
source .bashrc 
istioctl manifest apply --set profile=demo
