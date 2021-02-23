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

# 搜索 dockerfile 文件
dst="*dockerfile"
FS="$(find -name "$dst")"
echo "--> files: ${FS}"

# 处理文件

# 1 /eApi => /eapi
old="\/e[A-Z].*"
sed -n "/${old}/p" ${FS}
sed -i "s/\(${old}\)/\L\1/" ${FS}

# 2 delete goms-
sed -i "s/goms-//g" ${FS}

# 3 delete "apt-get update -y"
old="apt-get update -y"
sed -i "/${old}$/d" ${FS}

# 4
old="\/configs\/"
new="\/configs"
sed -i "s/${old}/${new}/g" ${FS}
