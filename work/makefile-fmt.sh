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
dst="*makefile"
files="$(find "$pro" -name "$dst")"
echo "--> files: $files"

file_set=("$files")
count=${#file_set[*]}
echo "--> count: $count"

# 处理文件
for file in $files; do
	# 添加 CHUB
	old='IMAGE\:=\$(NAME_SPACE)\/\$(REPO)\:\$(TAG)'
	new='IMAGE\:=\$(CHUB)\/\$(NAME_SPACE)\/\$(REPO)\:\$(TAG)'
	sed -i "s/$old/$new/g" "$file"
done
