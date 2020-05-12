#!/bin/bash
set -xe

# 版本信息
DASHBOARD_VERSION=v1.8.3

# pull 
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kubernetes-dashboard-amd64:$DASHBOARD_VERSION

# tag 带 amd ===> OK
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kubernetes-dashboard-amd64:$DASHBOARD_VERSION k8s.gcr.io/kubernetes-dashboard-amd64:$DASHBOARD_VERSION

# tag 不带 amd => ok
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kubernetes-dashboard-amd64:$DASHBOARD_VERSION k8s.gcr.io/kubernetes-dashboard:$DASHBOARD_VERSION

# rmi
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/kubernetes-dashboard-amd64:$DASHBOARD_VERSION


