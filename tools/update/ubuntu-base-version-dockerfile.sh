#!/bin/bash
set -x
set -e

# 当前目录路径
WD="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "--> work dir:${WD}"

# 当前项目路径
PD="$(cd "${WD}/../.." && pwd)"
echo "--> pro dir:${PD}"

# 搜索 dockerfile 文件
dst="*dockerfile"
files="$(cd ${PD} && find -name "$dst")"
echo "--> files: $files"

# work
old="FROM ubuntu:18.04"
new="FROM ubuntu:20.04"
(cd ${PD}; sed -i "s/${old}/${new}/g" ${files})