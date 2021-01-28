#!/bin/bash
# set -x
# set -e

# echo -e "==> start check ..."
echo -e "\033[34m==> start check ...\033[0m"

# 当前目录路径
pwdx="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# 当前项目路径 pro
pro="$(cd "$pwdx/../.." && pwd)"

# 工具目录 toolx
toolx=$pro/tools/hooks

# toolx_name 当前目录,工具集，不格式化
toolx_name=$(basename "$toolx")

# 为项目中的 bash 文加上运行权限
find "$pro" -name "*.sh" | xargs chmod +x

# cmd 获取改动的文件
files=
function changes_files() {
	wfiles="$(git diff --name-status | grep -v "^D" | awk '{ print $2; }')"
	cfiles="$(git diff --name-status --cached | grep -v "^D" | awk '{ print $2; }')"
	files="$wfiles $cfiles"
	# echo "$files"
}

# cmd 获取全部的文件，排除.git 和 工具目录
function all_files() {
	files="$(find $pro -name "*" -type f | grep -v .git | grep -v $toolx_name)"
	# echo "$files"
}

# 文档排版检查
[ "$1" ] || changes_files
[ "$1" = "all" ] && all_files
echo "--> files: $files"

# 文档排版检查
cd $pro
$toolx/check-doc.sh $files

# go 代码静态检查
files_go=$(echo "$files" | grep .go)
echo "--> files go: $files_go"
if [ -n "$files_go" ]; then
	cd $pro
	$toolx/check-code-go.sh $files_go
fi

# bash 代码静态检查
files_bash=$(echo "$files" | grep .sh)
echo "--> files bash: $files_bash"
if [ -n "$files_bash" ]; then
	cd $pro
	$toolx/check-code-bash.sh $files_bash
fi

# echo -e "==< end check"
echo -e "\033[34m==< end check\033[0m"
