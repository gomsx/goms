#!/bin/bash
set -xe

# 版本
# FLANNEL_VERSION=v0.10.0-amd64
FLANNEL_VERSION=v0.12.0-amd64

docker pull quay.io/coreos/flannel:$FLANNEL_VERSION