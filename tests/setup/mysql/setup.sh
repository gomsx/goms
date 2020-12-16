#!/bin/bash
set -e

PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD

echo ' ===> 1.启动mysql...'
service mysql start
sleep 3

echo `service mysql status`

echo ' ===> 2.创建库和表...'
bash $PWD/setup-data.sh
echo '导入完毕...'

echo `service mysql status`
echo ' ===> 3.初始化完成...'

tail -f /dev/null

