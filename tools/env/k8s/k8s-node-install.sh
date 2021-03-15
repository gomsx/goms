#!/bin/bash
set -x
set -e
set -u

## images
./images/k8s-pull-node.sh
./images/k8s-pull-flannel.sh

## env
./env/k8s-env-check.sh
./env/k8s-env-install.sh
./env/k8s-env-config.sh
./env/k8s-env-check.sh

## kubeadm
./kubeadm/node/kubeadm-install-node.sh
