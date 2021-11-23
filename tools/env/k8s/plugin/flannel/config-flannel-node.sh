#!/bin/bash
set -x

sysctl net.bridge.bridge-nf-call-iptables=1

# https://github.com/containernetworking/cni.dev
# https://github.com/flannel-io/flannel/blob/master/Documentation/running.md

rm -rf /run/flannel/
mkdir /run/flannel/

touch /run/flannel/subnet.env
chmod a+rwx /run/flannel/subnet.env

# ok
# 不同的 node FLANNEL_SUBNET 配置成不同的子网
cat <<EOF >/run/flannel/subnet.env
FLANNEL_NETWORK=10.1.0.0/16
FLANNEL_SUBNET=10.1.17.1/24
FLANNEL_MTU=1472
FLANNEL_IPMASQ=true
EOF
