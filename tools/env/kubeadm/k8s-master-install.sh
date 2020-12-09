#!/bin/bash
set -x
set -e
set -u

## env
./env/k8s-env-check.sh
./env/k8s-env-install.sh
./env/k8s-env-config.sh
./env/k8s-env-check.sh

## kubeadm
./kubeadm/master/kubeadm-install-master.sh

