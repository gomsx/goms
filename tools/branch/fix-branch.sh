#!/bin/bash
set -x
set -e

# branch 分支名
branch="$1"

FCMD="格式: bash_cmd branch"
ECMD="例子: fix-branch.sh dev"
HELP="❗ 分支名为空"

# 如果没有参数，打印提示并退出
if [ -z "${branch}" ]; then
	echo "${HELP}"
	echo "${FCMD}"
	echo "${ECMD}"

	branch="$(git symbolic-ref --short -q HEAD)"
	echo "输入的分支名为空, 使用当前分支: ${branch}"
else
	echo "输入的分支名 ${branch}"
fi

# 工作目录路径
WD="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# 项目目录路径
PD="$(cd "$WD/../.." && pwd)"

# README
## badge
sub="\/badge.svg?branch="
old="[^)]*)"
new="${branch})"
(cd ${PD} && sed -i "s/\(${sub}\)\(${old}\)/\1${new}/g" ./README.md)

## branch tree
old="tree\/[^\/]*\/"
new="tree\/${branch}\/"
(cd ${PD} && sed -i "s/${old}/${new}/g" ./README.md)
