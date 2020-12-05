#!/bin/bash
# set -x
set -e

# VER 版本参数
VER="$1"

# 如果没有参数，打印提示并退出
if [ -z $VER ];then
    echo -e "❗ 错误，缺版本参数 \n格式: bash_cmd src_dir \n例子：fix-branch.sh v1.0.1"
    exit
else
    echo "版本参数 $VER"
fi

# 当前 bash 所在目录路径 PWD
PWD=$(cd "$(dirname "$0")";pwd)

# 当前项目路径 PRO
PRO=$PWD/../..
PRO=$(cd $PRO;pwd)

# 替换版本参数
sed -i "s/_master/_$VER/g" $PRO/README.md
sed -i "s/\/master/\/release-$VER/g" $PRO/README.md

sed -i "s/_master/_$VER/g" $PRO/.github/workflows/make_master.yml
sed -i "s/master/release-$VER/g" $PRO/.github/workflows/make_master.yml
mv $PRO/.github/workflows/make_master.yml $PRO/.github/workflows/make_$VER.yml

