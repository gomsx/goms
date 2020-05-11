#!/bin/bash

set -x
set -e
set -u

# images
./images/k8s_pull_node.sh

# env
./env/k8s_env_check.sh
./env/k8s_env_install.sh
./env/k8s_env_config.sh
./env/k8s_env_check.sh

# kubeadm
./kubeadm/node/kubeadm_install_node.sh

