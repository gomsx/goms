#!/bin/bash

set -x
set -e
set -u

# PWD=$(cd "$(dirname "$0")";pwd)
# echo $PWD

# DK=$PWD/../docker
# K8S=$PWD

# $DK/install_docker.sh
# $DK/config_docker.sh

# env
./env/k8s_env_check.sh
./env/k8s_env_install.sh
./env/k8s_env_config.sh
./env/k8s_env_check.sh

# kubeadm
./kubeadm/node/kubeadm_install_node.sh



