#!/bin/bash
set -x
set -e

## env
bash env/k8s-env-check.sh
bash env/k8s-env-install.sh
bash env/k8s-env-config.sh
bash env/k8s-env-check.sh

## kubeadm
bash kubeadm/node/kubeadm-install-node.sh
