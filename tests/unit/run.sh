#!/bin/bash
set -x
set -e

# 当前目录路径
pwdx="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "==> pwdx:$pwdx"

# 当前项目路径 pro
prox="$(cd "$pwdx/../.." && pwd)"
echo "==> pro:$prox"

cd $prox/pkg;go test -v -gcflags=-l -count=1 ./...
cd $prox/eApi/build;make test
cd $prox/eTest/build;make test
