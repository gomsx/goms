#!/bin/bash
# set -x
set -e

echo -e "==> start check code bash ..."

# bash 源码静态分析
cd "$1"
shfmt $2

echo -e "==< end check code bash"
