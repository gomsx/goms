#!/bin/bash
set -x
set -e

echo -e "==> start xcopy ..."

# cp_dst 要复制的目标
cp_dst="$1"

# 如果没有目标参数，打印提示并退出
if [ -z $cp_dst ];then
    echo -e "❗ 错误，缺少目标参数 \n格式: bash_cmd copy_dst"
    exit
else
    echo "复制目标 $cp_dst"
fi

# 当前 bash 所在目录路径 PWD
PWD=$(cd "$(dirname "$0")";pwd)

# 当前项目路径 PRO
PRO=$PWD/../..
PRO=$(cd $PRO;pwd)

# DST 目标，SRC 源头
DST=$PRO/$cp_dst
SRC=${DST/fuwensun/vuca}

# 执行 copy
rm -rf $DST
cp -r $SRC $DST

$PWD/xcheck.sh

echo -e "==< end xcopy"

