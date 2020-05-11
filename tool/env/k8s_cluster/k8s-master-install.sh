#!/bin/bash

set -x
set -e
set -u

# images
./images/k8s_pull_master.sh
./images/k8s_pull_flannel.sh

# env
./env/k8s_env_check.sh
./env/k8s_env_install.sh
./env/k8s_env_config.sh
./env/k8s_env_check.sh

# kubeadm
./kubeadm/master/kubeadm_install_master.sh
./kubeadm/master/kubeadm_init_master.sh
./kubeadm/master/kubeadm_set_master.sh
./kubeadm/master/kubeadm_check.sh

# net
./kubeadm/plugin/config_flannel.sh
