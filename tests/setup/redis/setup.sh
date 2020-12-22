#!/bin/bash
set -x
set -e

# 当前目录路径
pwdx=$(
	cd "$(dirname "$0")"
	pwd
)

# 打印环境变量
whoami
pwd
which redis-server

# 启动 redis 服务
redis-server $pwdx/redis.conf

