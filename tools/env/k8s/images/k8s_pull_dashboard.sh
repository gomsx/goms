#!/bin/bash
set -xe

# 版本
DASHBOARD_VERSION=v2.0.0

docker pull kubernetesui/dashboard:$DASHBOARD_VERSION

