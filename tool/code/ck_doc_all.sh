#!/bin/bash
set -x

# PWD 
PWD=$(cd "$(dirname "$0")";pwd)
# PRO
PRO=$PWD/../..
PRO=$(cd $PRO;pwd)
BN=$(basename $PWD)
# CMD
CMD="find $PRO -name \"*\" -type f | grep -v .git | grep -v $BN" 
# FILES
FILES=$(eval $CMD)
echo "==> FILES: $FILES"

for f in $FILES
do
# 删除首行的空行
sed -i '/./,$!d' $f
# 匹配空格、tab等特殊字符
sed -i 's/^\s*$/\n/g' $f
# 尾行部插入空行
sed -i '$a\\n' $f
# 合并多个空行
sed -i '/^$/{N;/^\n*$/D}' $f
done

