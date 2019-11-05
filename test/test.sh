#!/bin/bash
set -e
set -x
set -u

# .sh所在目录
DIR=$(cd "$(dirname "$0")";pwd)
echo $DIR
PRODIR=$(cd "$DIR"/../;pwd)
echo $PRODIR;   

# 查看环境变量
echo $HOME; echo $PATH; which go;   
  
cd $PRODIR; go test ./...   # all
