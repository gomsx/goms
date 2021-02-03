#!/bin/bash
set -x
# set -e

# 当前目录路径
pwdx="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "==> pwdx:$pwdx"

# 当前项目路径 pro
pro="$(cd "$pwdx/.." && pwd)"
echo "==> pro:$pro"

# 搜索 makefile 文件
files="$(cd $pro;find . -name "*" -type f | grep -v .git \
| grep -v tools | grep -v work | grep -v eK8s | grep -v eIstio)"
echo "--> files: $files"

file_set=("$files")
count=${#file_set[*]}
echo "--> count: $count"

# exit

# 处理文件
for file in $files; do

	# 1 delete goms-
	sed -i "s/goms-//g" "$file"
done
