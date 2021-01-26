#!/bin/bash
# set -x
set -e

echo -e "==> start xcheck ..."

# 当前目录路径
pwdx="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# 当前项目路径
pro="$(cd "$pwdx/../.." && pwd)"

# 工具目录 toolx
toolx=$pro/tools/hooks

# toolx_name 当前目录,工具集，不格式化
toolx_name=$(basename "$toolx")

# 用 new 替换 old
old=fuwensun
new=aivuca

# cmd 搜索包含 old 的文件
cmd="grep $old -rl $pro --exclude-dir={.git,$toolx_name}"
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
