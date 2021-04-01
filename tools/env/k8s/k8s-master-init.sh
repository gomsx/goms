#!/bin/bash
set -x
set -e

## kubeadm
./kubeadm/master/kubeadm-init-master.sh
./kubeadm/master/kubeadm-set-master.sh
./kubeadm/master/kubeadm-check.sh
