#!/bin/bash

set -x
set -e
set -u

## kubeadm
./kubeadm/node/kubeadm_join_node.sh

