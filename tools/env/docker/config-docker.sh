#!/bin/bash
# Ubuntu18.04,Ubuntu20.04
set -x
set -e
set -u

sudo mkdir -p /etc/docker
sudo tee /etc/docker/daemon.json <<-'EOF'
{
    "registry-mirrors": [
    "https://registry.cn-hangzhou.aliyuncs.com",
    "https://registry.docker-cn.com",
    "http://hub-mirror.c.163.com"
    ],
    "dns": ["8.8.8.8","8.8.4.4"]
}
EOF

sudo systemctl daemon-reload
sudo systemctl restart docker
