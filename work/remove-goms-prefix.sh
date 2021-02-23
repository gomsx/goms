#!/bin/bash
set -x
set -e

# 工作目录路径
WD="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "--> work dir: ${WD}"

# 项目目录路径
PD="$(cd "$WD/.." && pwd)"
echo "--> pro dir: ${PD}"

# cd 
cd ${PD}

# 搜索文件
dst="goms-"
FS="$(grep -rl "${dst}" --exclude-dir=.git \
| grep -v tools | grep -v work | grep -v eK8s | grep -v eIstio)"
echo "--> files: ${FS}"

# work
sed -i "s/${dst}//g" ${FS}
