#!/bin/bash
set -x
set -e

# 当前目录路径
pwdx=$(
	cd "$(dirname "$0")"
	pwd
)

service redis stop

redis-server "$pwdx"/redis.conf

