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
	# 2 := -> =
	old=":="
	new="="
	sed -i "s/$old/$new/g" "$file"

	# 3 go test
	old="go test \.\.\/\.\.\. -v -gcflags=-l -count=1"
	new="go test -v -gcflags=-l -count=1 \.\.\/\.\.\."
	sed -i "s/$old/$new/g" "$file"

	# 4 eApi => eapi
	old="=e[A-Z].*"
	# sed -n "/$old/p" "$file"
	sed -i "s/\($old\)/\L\1/" "$file"

	# 5 delete REPO=goms-eapi
	old="REPO="
	sed -i "/APP=/{ n;/$old/d }" "$file"
	sed -i "s/goms-//g" "$file"

	# 6 REPO => APP
	old="REPO"
	new="APP"
	sed -i "s/$old/$new/g" "$file"
done
