#!/bin/bash
set -x
set -e
set -u

## kubeadm
./kubeadm/node/kubeadm-join-node.sh

