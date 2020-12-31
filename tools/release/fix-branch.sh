#!/bin/bash
# set -x
set -e

# ver 版本参数
ver="$1"

# 如果没有参数，打印提示并退出
if [ -z $ver ];then
    echo -e "❗ 错误，缺版本参数 \n格式: bash_cmd src_dir \n例子：fix-branch.sh v1.0.1"
    exit
else
    echo "版本参数 $ver"
fi

# 当前目录路径
pwdx=$(
	cd "$(dirname "$0")"
	pwd
)

# 当前项目路径 pro
pro=$(
	cd "$pwdx/../.."
	pwd
)

# workflow 目录
wfd=$pro/workflows

# 替换版本参数
sed -i "s/_main/_$ver/g" $pro/README.md
sed -i "s/\/main/\/release-$ver/g" $pro/README.md

sed -i "s/_main/_$ver/g" $wfd/make_main.yml
sed -i "s/main/release-$ver/g" $wfd/make_main.yml
mv $wfd/make_main.yml $wfd/make_$ver.yml

