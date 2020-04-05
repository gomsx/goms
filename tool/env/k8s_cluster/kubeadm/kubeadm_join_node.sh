#!/bin/bash
set -xe

sudo kubeadm reset
sudo kubeadm join 192.168.43.201:6443 --token vpglmp.402txzbv97s1lf7r --discovery-token-ca-cert-hash sha256:2ffbec1590c1a97e122035762fff6bbffe8272c01d618991d7880ac4e3eee3f6
