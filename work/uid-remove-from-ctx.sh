#!/bin/bash
set -x
# set -e

# 当前目录路径
pwdx="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "==> pwdx:$pwdx"

# 当前项目路径 pro
pro="$(cd "$pwdx/.." && pwd)"
echo "==> pro:$pro"

# 搜索 go 文件
dst="*.go"
files="$(find "$pro" -name "$dst")"
echo "--> files: $files"

file_set=("$files")
count=${#file_set[*]}
echo "--> count: $count"

# 处理文件
for file in $files; do

	# delete
	old="ms.CarryCtxUserId"
	sed -i "/$old/d" "$file"

	old="ms.GetCtxVal(ctx)"
	sed -i "/$old/d" "$file"

	# c -> ctx
	old="(c context"
	new="(ctx context"
	sed -i "s/$old/$new/g" "$file"

	old="(c,"
	new="(ctx,"
	sed -i "s/$old/$new/g" "$file"

	old="(c)"
	new="(ctx)"
	sed -i "s/$old/$new/g" "$file"

	# patch
	old="Notify(ctx,"
	new="Notify(c,"
	sed -i "s/$old/$new/g" "$file"

done
