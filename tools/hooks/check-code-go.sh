#!/bin/bash
# set -x
# set -e

echo -e "==> start check code go ..."

files="$1"

pkgs=$(echo "${files}" | xargs dirname | sort -u)

for pkg in ${pkgs}; do
    go fmt ./${pkg}/...
    go vet ./${pkg}/...
done

# go mod tidy

# 逃逸分析，'-l' 防止内联
# go build -gcflags '-m -l'

echo -e "==< end check code go"
