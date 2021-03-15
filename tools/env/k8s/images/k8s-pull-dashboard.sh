#!/bin/bash
set -x
set -e

# 版本
dashboard_version=v2.0.0

docker pull kubernetesui/dashboard:$dashboard_version
