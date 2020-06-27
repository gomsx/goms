#!/bin/bash
set -xe

sudo kubeadm reset
sudo kubeadm join 192.168.43.201:6443 --token topaok.9xudyler0jca3qld --discovery-token-ca-cert-hash sha256:915c6bb06c3b581b23b5c
