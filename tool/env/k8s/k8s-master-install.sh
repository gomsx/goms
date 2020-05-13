#!/bin/bash

set -x
set -e
set -u

## env
./env/k8s_env_check.sh
./env/k8s_env_install.sh
./env/k8s_env_config.sh
./env/k8s_env_check.sh

## kubeadm
./kubeadm/master/kubeadm_install_master.sh