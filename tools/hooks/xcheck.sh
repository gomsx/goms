#!/bin/bash
set -x
set -e

echo -e "==> start xcheck ..."

# 当前 bash 所在目录路径 pwdx
pwdx=$(
	cd "$(dirname "$0")"
	pwd
)

# 当前项目路径 pro
pro=$pwdx/../..
pro=$(
	cd "$pro"
	pwd
)

# pwdx_name 当前目录,工具集，不格式化
pwdx_name=$(basename "$pwdx")

# 用 new 替换 old
old=fuwensun
new=aivuca

# cmd 搜索包含 old 的文件
cmd="grep $old -rl $pro --exclude-dir={.git,$pwdx_name}"
cmde="grep $old -rl $pro"

# files 包含 old 的文件集合
echo "替换前："
files=$(eval "$cmd")
echo "--> files: $files"

# files count 文件数
file_set=("$files")
count=${#file_set[*]}
echo "--> count: $count"

# 执行替换
echo "替换："
sed -i "s/$old/$new/g" $files # $files ==> f1 f2 ... # "$files" ==> 'f1 f2'

echo "替换后："
eval "$cmde"

echo -e "==< end xcheck"
