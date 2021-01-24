#!/bin/bash
# set -x
set -e

# dependce 依赖组件
dep="$1"
ver="$2"

# 如果没有参数，打印提示并退出
fcmd="bash_cmd dep ver"
ecmd="dependence-version.sh mysqltest versoin"
help="❗ 依赖组件 \n格式: $fcmd\n例子: $ecmd"

# dep
if [ -z $dep ]; then
	echo -e "$help"
	exit
else
	echo -e "依赖组件 $dep"
fi

# ver
if [ -z $ver ]; then
	echo -e "$help"
	exit
else
	echo "依赖版本 $ver" # TODO 自动搜索 tests/*
fi

# 当前目录路径
pwdx="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "==> pwdx:$pwdx"

# 当前项目路径 pro
pro="$(cd "$pwdx/../.." && pwd)"
echo "==> pro:$pro"
# exit

# 目标目录
dir_ary=("$pro/.github/workflows" "$pro/eApi/internal/dao" "$pro/eTest/internal/dao")
dirs="${dir_ary[@]}"
echo "==> dirs:$dirs"
# exit

i=0
for dir in $dirs; do
	files[i++]="$(grep -rl "$dep" "$dir")"
done
files="${files[@]}"
echo "==> files:$files"
# exit

for file in $files; do
	# sed -n "/"$dep:".*$/p" "$file"
	sed -i "s/"$dep:".*$/"$dep"\:"$ver"/g" "$file"
done
