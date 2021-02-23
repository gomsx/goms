#!/bin/bash
set -x
set -e

# 工作目录路径
WD="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "--> work dir: ${WD}"

# 项目目录路径
PD="$(cd "$WD/.." && pwd)"
echo "--> pro dir: ${PD}"

cd ${PD}

# 搜索 makefile 文件
dst="*makefile"
FS="$(cd ${PD} && find -name "$dst")"
echo "--> files: ${FS}"

old="go test -v"
new="go test"
sed -i "s/${old}/${new}/g" ${FS}

# 搜索 sh 文件
dst="*.sh"
FS="$(cd ${PD} && find -name "$dst" | grep test)"
echo "--> files: ${FS}"

old="go test -v"
new="go test"
sed -i "s/${old}/${new}/g" ${FS}
