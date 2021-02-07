#!/bin/bash
set -x
set -e

# 项目 root
prox="$(cd ../../ && pwd)"
echo "==> prox:$prox"

cd ${prox}/pkg; go test -v -gcflags=-l -count=1 ./...
cd ${prox}/eApi/build; make test
cd ${prox}/eTest/build; make test
