#!/bin/bash
set -x
set -e

# get version
kubeadm version
kubectl version
kubelet --version
