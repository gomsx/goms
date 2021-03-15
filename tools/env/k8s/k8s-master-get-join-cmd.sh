#!/bin/bash
set -x
set -e
set -u

./kubeadm/master/kubeadm-create-token.sh
./kubeadm/master/kubeadm-gen-join-cmd.sh
