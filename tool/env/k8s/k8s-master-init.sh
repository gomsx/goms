#!/bin/bash

set -x
set -e
set -u

## images
./images/k8s_pull_master.sh
./images/k8s_pull_flannel.sh

## kubeadm
./kubeadm/master/kubeadm_init_master.sh
./kubeadm/master/kubeadm_set_master.sh
./kubeadm/master/kubeadm_check.sh

## addons
### net
./plugin/flannel/config_flannel.sh
./plugin/flannel/deploy_flannel.sh

### dashboard
./plugin/dashboard/deploy_dashboard.sh

