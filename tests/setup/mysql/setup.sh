#!/bin/bash
set -x
# set -e

# 当前目录路径
pwdx=$(
	cd "$(dirname "$0")"
	pwd
)

echo ' ===> 1.启动 mysql...'
service mysql start
sleep 5s

service mysql status

echo ' ===> 2.创建库和表...'
bash "$pwdx"/setup-data.sh
echo '导入完毕...'

service mysql status
echo ' ===> 3.初始化完成...'

tail -f /dev/null
