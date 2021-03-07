#!/bin/bash
set -x
set -e

# 项目目录
PD="$(cd ../../ && pwd)"
echo "==> pro dir: ${PD}"

cd ${PD}/pkg; go test -gcflags=-l -count=1 ./...
cd ${PD}/eApi/build; make test
cd ${PD}/eTest/build; make test
