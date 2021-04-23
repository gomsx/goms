#!/bin/bash
set -x
set -e

## env
bash env/k8s-env-check.sh
bash env/k8s-env-install.sh
bash env/k8s-env-config.sh
bash env/k8s-env-check.sh

## kubeadm
bash kubeadm/master/kubeadm-install-master.sh

## version
bash kubeadm/master/kubeadm-check.sh
