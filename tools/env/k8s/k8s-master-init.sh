#!/bin/bash
set -x
set -e

## kubeadm
bash kubeadm/master/kubeadm-init-master.sh
bash kubeadm/master/kubeadm-set-master.sh
bash kubeadm/master/kubeadm-check.sh
