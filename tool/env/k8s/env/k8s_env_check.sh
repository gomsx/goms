#!/bin/bash

# set -x
# set -e
# set -u

succ="[succ]"
failed="[failed]"

# 关闭 selinux
# https://blog.csdn.net/wx740851326/article/details/72302931
res=$(getenforce)
echo "==> selinux $res"
[[ $res = "Disabled" ]] && echo $succ || echo $failed

# 关闭 swap
# https://blog.csdn.net/ygm_linux/article/details/24532809
res=$(free xargs | awk 'NR==3{ print $2 $3 $4 }')
echo "==> swap $res"
[[ $res = "000" ]] && echo $succ || echo $failed

# 关闭 ufw 防火墙
# https://blog.csdn.net/liukuan73/article/details/83116271
res=$(sudo ufw status)
#去空格
res="$(echo -e "${res}" | tr -d '[:space:]')"
res=${res##*Status:}
echo "==> ufw $res"
[[ $res = "inactive" ]] && echo $succ || echo $failed