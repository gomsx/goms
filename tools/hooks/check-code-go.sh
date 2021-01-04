#!/bin/bash
# set -x
set -e

echo -e "==> start check code go ..."

# go 源码静态分析
cd "$1"
for f in $2; do
    go fmt $f
done
go mod tidy

# 逃逸分析，'-l' 防止内联
# go build -gcflags '-m -l'

echo -e "==< end check code go"
