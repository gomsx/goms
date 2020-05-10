#!/bin/bash

# set -x
# set -e
# set -u

# Selinux
# https://blog.csdn.net/wx740851326/article/details/72302931
getenforce
echo " ==> 需要 Disabled"

# swap
# https://blog.csdn.net/ygm_linux/article/details/24532809
free
echo " ==> 需要 0 0 0 0"

# ufw防火墙
# https://blog.csdn.net/liukuan73/article/details/83116271
sudo ufw status