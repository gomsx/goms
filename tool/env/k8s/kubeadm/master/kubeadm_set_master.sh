#!/bin/bash
set -xe

rm -rf $HOME/.kube/*

ug=$(id -u):$(id -g)
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $ug $HOME/.kube/config
