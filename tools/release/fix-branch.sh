#!/bin/bash
# set -x
set -e

# version 版本参数
version="$1"

# 如果没有参数，打印提示并退出
if [ -z $version ]; then
	echo -e "❗ 错误，缺版本参数 \n格式: bash_cmd version \n例子：fix-branch.sh v1.0.1"
	exit
else
	echo "版本参数 $version"
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
workflow=$pro/.github/workflows

# 替换版本参数
sed -i "s/make_main/make_$version/g" $pro/README.md
sed -i "s/\/tree\/main/\/tree\/release-$version/g" $pro/README.md

sed -i "s/make_main/make_$version/g" $workflow/make_main.yml
sed -i "s/\[ main \]/\[ release-$version \]/g" $workflow/make_main.yml
mv $workflow/make_main.yml $workflow/make_$version.yml
