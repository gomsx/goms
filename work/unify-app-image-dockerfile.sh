#!/bin/bash
set -x
# set -e

# 当前目录路径
pwdx="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "==> pwdx:$pwdx"

# 当前项目路径 pro
pro="$(cd "$pwdx/.." && pwd)"
echo "==> pro:$pro"

# 搜索 dockerfile 文件
dst="*dockerfile"
files="$(find "$pro" -name "$dst")"
echo "--> files: $files"

file_set=("$files")
count=${#file_set[*]}
echo "--> count: $count"

# 处理文件
for file in $files; do

	# 1 /eApi => /eapi
	old="\/e[A-Z].*"
	sed -n "/$old/p" "$file"
	sed -i "s/\($old\)/\L\1/" "$file"

	# 2 delete goms-
	sed -i "s/goms-//g" "$file"

	# 3 delete "apt-get update -y"
	old="apt-get update -y"
	sed -i "/$old$/d" "$file"

	# 4
	sed -i "s/\/configs\//\/configs/g" "$file"
done
