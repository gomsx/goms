#!/bin/bash
set -x
set -e

# branch 分支名
branch="$1"

FCMD="bash_cmd branch"
ECMD="fix-branch.sh dev"
HELP="❗ 分支名 \n格式: ${FCMD}\n例子: ${ECMD}"

# 如果没有参数，打印提示并退出
if [ -z "${branch}" ]; then
	echo -e "${HELP}"
	exit
else
	echo "分支名 ${branch}"
fi

# 项目 root
cd ../../
pwd

cicd_name="cicd-${branch}"
cicd_file="cicd.yml"

# README
## badge
old="workflows\/[^\/]*\/badge"
new="workflows\/${cicd_name}\/badge"
sed -i "s/${old}/${new}/g" ./README.md

old="workflow%3A.*)"
new="workflow%3A${cicd_name})"
sed -i "s/${old}/${new}/g" ./README.md

## branch tree
old="tree\/[^\/]*\/"
new="tree\/${branch}\/"
sed -i "s/${old}/${new}/g" ./README.md

# workflow 目录
cd .github/workflows
pwd

## cicd
cicd="cicd.yml"
old="^name\: .*$"
new="name\: ${cicd_name}"
sed -i "s/${old}/${new}/g" ${cicd_file}

## .yml
old="branches: \[ .* \]"
new="branches: \[ ${branch} \]"
sed -i "s/${old}/${new}/g" "$(grep -rl -E "${old}")"
