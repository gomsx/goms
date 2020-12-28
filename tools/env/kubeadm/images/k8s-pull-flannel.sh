#!/bin/bash
set -x
set -e

# 版本
flannel_version=v0.12.0-amd64

docker pull quay.io/coreos/flannel:$flannel_version

