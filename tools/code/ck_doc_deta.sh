#!/bin/bash
set -x

# PWD 
PWD=$(cd "$(dirname "$0")";pwd)
# PRO
PRO=$PWD/../..
PRO=$(cd $PRO;pwd)
BN=$(basename $PWD)
# CMD
CMD="git status -s | awk '{ print \$2; }'" # $2 要做字符串处理，即 \$2
# FILES
FILES=$(eval $CMD)
echo "==> FILES: $FILES"

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

