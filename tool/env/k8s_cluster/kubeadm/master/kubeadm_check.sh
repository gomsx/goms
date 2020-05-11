#!/bin/bash
set -xe

kubeadm version
kubectl version
kubelet --version

# kubeadm token create
# openssl x509 -pubkey -in /etc/kubernetes/pki/ca.crt | openssl rsa -pubin -outform der 2>/dev/null | openssl dgst -sha256 -hex | sed 's/^.* //'
