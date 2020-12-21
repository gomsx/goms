#!/bin/bash
# set -x
set -e

echo -e "==> start check code go ..."

# 当前目录路径
pwdx=$(
	cd "$(dirname "$0")"
	pwd
)

# 当前项目路径 pro
pro=$(
	cd "$pwdx/../.."
	pwd
)

# go 源码静态分析
cd "$pro"
go fmt ./...
go mod tidy

# 逃逸分析，'-l' 防止内联
# go build -gcflags '-m -l'

echo -e "==< end check code go"
