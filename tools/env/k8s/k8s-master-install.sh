#!/bin/bash
set -x
set -e

## env
./env/k8s-env-check.sh
./env/k8s-env-install.sh
./env/k8s-env-config.sh
./env/k8s-env-check.sh

## kubeadm
./kubeadm/master/kubeadm-install-master.sh

## version
./kubeadm/master/kubeadm-check.sh
