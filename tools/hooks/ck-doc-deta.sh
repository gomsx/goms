#!/bin/bash
# set -x
set -e

echo -e "==> start check doc deta ..."

# 当前 bash 所在目录路径 PWD
PWD=$(cd "$(dirname "$0")";pwd)

# 当前项目路径 PRO
PRO=$PWD/../..
PRO=$(cd $PRO;pwd)

# PWD_NAME 当前目录,工具集，不格式化
PWD_NAME=$(basename $PWD)

# CMD 获取改动的文件
CMD="git status -s | awk '{ print \$2; }' | grep -v /$PWD_NAME" # $2 要做字符串处理，即 \$2

# FILES 改动的集合
FILES=$(eval $CMD)
echo "--> FILES: $FILES"

# 处理改动的文件
for f in $FILES
do
# 匹配空格、tab等特殊字符,替换成换行符
sed -i 's/^\s*$/\n/g' $f
# 尾行部插入空行
sed -i '$a\\n' $f
# 合并多个空行
sed -i '/^$/{N;/^\n*$/D}' $f
# 删除为空的首行
sed -i '/./,$!d' $f
done

echo -e "==< end check doc deta"

