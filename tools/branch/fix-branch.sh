#!/bin/bash
# set -x
set -e

# branch 分支名
branch="$1"

# 如果没有参数，打印提示并退出
if [ -z $branch ]; then
	echo -e "❗ 分支名 \n格式: bash_cmd branch \n例子：fix-branch.sh dev"
	exit
else
	echo "分支名 $branch"
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
sed -i "s/make_main/make_$branch/g" $pro/README.md
sed -i "s/\/tree\/main/\/tree\/$branch/g" $pro/README.md

sed -i "s/make_main/make_$branch/g" $workflow/make_main.yml
sed -i "s/\[ main \]/\[ $branch \]/g" $workflow/make_main.yml
mv $workflow/make_main.yml $workflow/make_$branch.yml
