#!/bin/bash
set -xe

# 版本信息
K8S_VERSION=v1.15.1
ETCD_VERSION=3.3.10

DASHBOARD_VERSION=v1.8.3
FLANNEL_VERSION=v0.10.0-amd64
DNS_VERSION=1.14.8

PAUSE_VERSION=3.1
COREDNS_VERSION=1.3.1

# pull
## 基本组件
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/kube-proxy-amd64:$K8S_VERSION
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/pause-amd64:$PAUSE_VERSION
## 网络
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/k8s-dns-sidecar-amd64:$DNS_VERSION
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/k8s-dns-kube-dns-amd64:$DNS_VERSION
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/k8s-dns-dnsmasq-nanny-amd64:$DNS_VERSION
docker pull quay.io/coreos/flannel:$FLANNEL_VERSION

# tag
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-proxy-amd64:$K8S_VERSION k8s.gcr.io/kube-proxy-amd64:$K8S_VERSION
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/pause-amd64:$PAUSE_VERSION k8s.gcr.io/pause-amd64:$PAUSE_VERSION

docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/k8s-dns-sidecar-amd64:$DNS_VERSION k8s.gcr.io/k8s-dns-sidecar-amd64:$DNS_VERSION
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/k8s-dns-kube-dns-amd64:$DNS_VERSION k8s.gcr.io/k8s-dns-kube-dns-amd64:$DNS_VERSION
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/k8s-dns-dnsmasq-nanny-amd64:$DNS_VERSION k8s.gcr.io/k8s-dns-dnsmasq-nanny-amd64:$DNS_VERSION

#docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/kube-proxy-amd64:$K8S_VERSION k8s.gcr.io/kube-proxy:$K8S_VERSION
#docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/pause-amd64:$PAUSE_VERSION k8s.gcr.io/pause:$PAUSE_VERSION
#docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/k8s-dns-sidecar-amd64:$DNS_VERSION k8s.gcr.io/k8s-dns-sidecar:$DNS_VERSION
#docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/k8s-dns-kube-dns-amd64:$DNS_VERSION k8s.gcr.io/k8s-dns-kube-dns:$DNS_VERSION
#docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/k8s-dns-dnsmasq-nanny-amd64:$DNS_VERSION k8s.gcr.io/k8s-dns-dnsmasq-nanny:$DNS_VERSION

## rmi
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/kube-proxy-amd64:$K8S_VERSION
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/pause-amd64:$PAUSE_VERSION

docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/k8s-dns-sidecar-amd64:$DNS_VERSION
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/k8s-dns-kube-dns-amd64:$DNS_VERSION
docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/k8s-dns-dnsmasq-nanny-amd64:$DNS_VERSION


