#!/bin/bash
set -x

# pwd
PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD

# PRO
PRO=$PWD/../..

# chmod *.sh
find $PRO -name "*.sh" | xargs chmod +x 

# go
go fmt $PRO/...

# 逃逸分析，'-l' 防止内联
# go build -gcflags '-m -l'

###############################
# PWD 
PWD=$(cd "$(dirname "$0")";pwd)
# PRO
PRO=$PWD/../..
PRO=$(cd $PRO;pwd)
BN=$(basename $PWD)
# CMD
CMD="grep * -rl $PRO --exclude-dir={.git,$BN}"
# FILES
FILES=$(eval $CMD)
echo "==> FILES: $FILES"

# 删除首行的空行
sed -i '/./,$!d' $FILES
# 尾行部插入空行
sed -i '$a\\n' $FILES
# 合并多个空行
sed -i '/^$/{N;/^\n*$/D}' $FILES
