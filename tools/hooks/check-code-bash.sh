#!/bin/bash
# set -x
# set -e

echo -e "==> start check code bash ..."

files="$1"

# bash 源码静态分析
shfmt -w ${files}

echo -e "==< end check code bash"
