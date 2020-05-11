#!/bin/bash
set -xe

# 版本信息
FLANNEL_VERSION=v0.10.0-amd64

docker pull quay.io/coreos/flannel:$FLANNEL_VERSION