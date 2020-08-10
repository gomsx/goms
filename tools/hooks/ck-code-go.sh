#!/bin/bash
# set -x

echo -e "==> start check code go ..."

# pwd
PWD=$(cd "$(dirname "$0")";pwd)
# echo "--> PWD: $PWD"

# PRO
PRO=$PWD/../..
# echo "--> PRO: $PRO"

# go fmt
cd $PRO;go fmt ./...;go mod tidy

# 逃逸分析，'-l' 防止内联
# go build -gcflags '-m -l'

echo -e "==> end check code go"

