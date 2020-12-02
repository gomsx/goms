#!/bin/bash
# set -x
set -e

echo -e "==> start xcheck ..."

# 当前 bash 所在目录路径 PWD
PWD=$(cd "$(dirname "$0")";pwd)

# 当前项目路径 PRO
PRO=$PWD/../..
PRO=$(cd $PRO;pwd)

# PWD_NAME 当前目录,工具集，不格式化
PWD_NAME=$(basename $PWD)

# 用 NEW 替换 OLD 
OLD=fuwensun
NEW=aivuca

# CMD 搜索包含 OLD 的文件
CMD="grep $OLD -rl $PRO --exclude-dir={.git,$PWD_NAME}"
CMDE="grep $OLD -rl $PRO"

# FILES 包含 OLD 的文件集合
echo "替换前："
FILES=$(eval $CMD)
echo "--> FILES: $FILES"

# FILES COUNT 文件数
FILE_SET=($FILES)
COUNT=${#FILE_SET[*]}
echo "--> COUNT: $COUNT"

# 执行替换
echo "替换："
sed -i "s/$OLD/$NEW/g"  $FILES

echo "替换后："
eval $CMDE

echo -e "==< end xcheck"

