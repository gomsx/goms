#!/bin/bash
set -x
set -e
set -u

## images
./images/k8s-pull-master.sh
./images/k8s-pull-flannel.sh

## kubeadm
./kubeadm/master/kubeadm-init-master.sh
./kubeadm/master/kubeadm-set-master.sh
./kubeadm/master/kubeadm-check.sh

## addons
### net
./plugin/flannel/config-flannel.sh
./plugin/flannel/deploy-flannel.sh

### dashboard
./plugin/dashboard/deploy-dashboard.sh

