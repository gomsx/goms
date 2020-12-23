#!/bin/bash
set -x
set -e

# 版本信息
k8s_version=v1.18.1
etcd_version=3.4.3-0
pause_version=3.2
coredns_version=1.6.7

# master 基本组件
## pull
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kube-apiserver-amd64:$k8s_version
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kube-controller-manager-amd64:$k8s_version
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kube-scheduler-amd64:$k8s_version
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/etcd-amd64:$etcd_version

## tag 带 amd
# docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-apiserver-amd64:$k8s_version k8s.gcr.io/kube-apiserver-amd64:$k8s_version
# docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-controller-manager-amd64:$k8s_version k8s.gcr.io/kube-controller-manager-amd64:$k8s_version
# docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-scheduler-amd64:$k8s_version k8s.gcr.io/kube-scheduler-amd64:$k8s_version
# docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/etcd-amd64:$etcd_version k8s.gcr.io/etcd-amd64:$etcd_version

## tag 无 amd => ok
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-apiserver-amd64:$k8s_version k8s.gcr.io/kube-apiserver:$k8s_version
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-controller-manager-amd64:$k8s_version k8s.gcr.io/kube-controller-manager:$k8s_version
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-scheduler-amd64:$k8s_version k8s.gcr.io/kube-scheduler:$k8s_version
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/etcd-amd64:$etcd_version k8s.gcr.io/etcd:$etcd_version

## rmi
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/kube-apiserver-amd64:$k8s_version
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/kube-controller-manager-amd64:$k8s_version
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/kube-scheduler-amd64:$k8s_version
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/etcd-amd64:$etcd_version

####################################################################################
## 其他
# pull ## 网络
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kube-proxy-amd64:$k8s_version
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/pause-amd64:$pause_version
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/coredns:$coredns_version

# tag 带 amd
# docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-proxy-amd64:$k8s_version k8s.gcr.io/kube-proxy-amd64:$k8s_version
# docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/pause-amd64:$pause_version k8s.gcr.io/pause-amd64:$pause_version
# docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/coredns:$coredns_version k8s.gcr.io/coredns:$coredns_version

# tag 不带 amd => ok
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-proxy-amd64:$k8s_version k8s.gcr.io/kube-proxy:$k8s_version
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/pause-amd64:$pause_version k8s.gcr.io/pause:$pause_version
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/coredns:$coredns_version k8s.gcr.io/coredns:$coredns_version

# rmi
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/kube-proxy-amd64:$k8s_version
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/pause-amd64:$pause_version
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/coredns:$coredns_version

