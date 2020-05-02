#!/bin/bash

set -x
set -e
set -u

sudo mkdir -p /etc/docker
sudo tee /etc/docker/daemon.json <<-'EOF'
{
    "registry-mirrors": [
    "https://kfwkfulq.mirror.aliyuncs.com",
    "https://2lqq34jg.mirror.aliyuncs.com",
    "https://pee6w651.mirror.aliyuncs.com",
    "https://registry.docker-cn.com",
    "http://hub-mirror.c.163.com"
    ],
    "dns": ["8.8.8.8","8.8.4.4"]
}
EOF

sudo systemctl daemon-reload
sudo systemctl restart docker