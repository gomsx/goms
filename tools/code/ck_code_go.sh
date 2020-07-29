#!/bin/bash
set -x

# pwd
PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD

# PRO
PRO=$PWD/../..

# go fmt
cd $PRO;go fmt ./...;go mod tidy

# 逃逸分析，'-l' 防止内联
# go build -gcflags '-m -l'

