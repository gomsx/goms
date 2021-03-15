#!/bin/bash
set -x
set -e

# 版本信息
k8s_version=v1.18.1
pause_version=3.2
coredns_version=1.6.7

# 基本组件
## pull
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kube-proxy-amd64:$k8s_version
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/pause-amd64:$pause_version
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/coredns:$coredns_version

## tag, 带 amd
# docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-proxy-amd64:$k8s_version k8s.gcr.io/kube-proxy-amd64:$k8s_version
# docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/pause-amd64:$pause_version k8s.gcr.io/pause-amd64:$pause_version

## tag, 不带 amd => ok
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-proxy-amd64:$k8s_version k8s.gcr.io/kube-proxy:$k8s_version
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/pause-amd64:$pause_version k8s.gcr.io/pause:$pause_version
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/coredns:$coredns_version k8s.gcr.io/coredns:$coredns_version

## rmi
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/kube-proxy-amd64:$k8s_version
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/pause-amd64:$pause_version
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/coredns:$coredns_version
