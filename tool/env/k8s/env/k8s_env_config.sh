#!/bin/bash
set -xe


# 关闭 selinux
# https://segmentfault.com/q/1010000018470347
# setenforce 0
# sed -i "s/SELINUX=enforcing/SELINUX=disabled/g" /etc/selinux/config

# 关闭 swap
# https://blog.csdn.net/CSDN_duomaomao/article/details/75142769
sudo swapoff -a
sudo mount -n -o remount,rw /

# 关闭 ufw 防火墙
# https://blog.csdn.net/liukuan73/article/details/83116271
sudo ufw disable

