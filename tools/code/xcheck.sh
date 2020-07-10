#!/bin/bash
# set -x

# PWD 
PWD=$(cd "$(dirname "$0")";pwd)
echo "==> PWD: $PWD"

# PRO
PRO=$PWD/../..
PRO=$(cd $PRO;pwd)
echo "==> PRO: $PRO"

# BN=${PWD##*"goms/"}
# BN=${PWD##"$PRO/"}
BN=$(basename $PWD)
echo "==> BN: $BN"

# replace
S1=fuwensun
# S1=waoops
S2=aivuca

# CMD
CMD="grep $S1 -rl $PRO --exclude-dir={.git,$BN}"
CMDX="grep $S1 -rl $PRO --exclude-dir={.git,$BN}"
CMDE="grep $S1 -rl $PRO"
echo "==> CMD: $CMD"

# FILES
FILES=$(eval $CMD)
echo "==> FILES: $FILES"

# FILES COUNT
arr=($FILES)
COUNT=${#arr[*]}
echo "==> COUNT: $COUNT"

# RES
echo "=============<CMDX-1>==============>"
echo "==> CMDX: $CMDX"
eval $CMDX

echo "=============<sed>================>"
sed -i "s/$S1/$S2/g"  $FILES

echo "=============<CMDX-2>==============>"
echo "==> CMDX: $CMDX"
eval $CMDX

echo "=============<CMDE>==============>"
echo "==> CMDE: $CMDE"
eval $CMDE