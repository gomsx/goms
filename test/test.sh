#!/bin/bash
set -e
set -x
set -u

# 工程目录
## .sh所在目录
DIR=$(cd "$(dirname "$0")";pwd)
echo $DIR
## 工程目录
PRODIR=$(cd "$DIR"/../;pwd)
echo $PRODIR;   

# 设置环境变量

## 查看
echo $HOME; 
echo $PATH; 
which go;   

# 测试
cd $PRODIR; go test ./...   # all
