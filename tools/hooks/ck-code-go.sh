#!/bin/bash
# set -x
set -e

echo -e "==> start check code go ..."

# 当前 bash 所在目录路径 PWD
PWD=$(cd "$(dirname "$0")";pwd)

# 当前项目路径 PRO
PRO=$PWD/../..
PRO=$(cd $PRO;pwd)

# go 源码静态分析
cd $PRO
go fmt ./...
go mod tidy

# 逃逸分析，'-l' 防止内联
# go build -gcflags '-m -l'

echo -e "==< end check code go"

