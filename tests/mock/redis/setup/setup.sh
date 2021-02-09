#!/bin/bash
set -x
set -e

# 启动 redis 服务
redis-server redis.conf
