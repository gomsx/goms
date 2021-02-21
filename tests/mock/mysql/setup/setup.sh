#!/bin/bash
set -x
set -e

# 设置数据库
echo ' ===> 1.启动 mysql...'
service mysql start
# sleep 5s

service mysql status

echo ' ===> 2.创建库和表...'
bash setup-data.sh
echo '导入完毕...'

service mysql status
echo ' ===> 3.初始化完成...'
