#!/bin/bash
set -x
set -e

# dependce 依赖组件
dep="$1"
ver="$2"

# 如果没有参数，打印提示并退出
FCMD="bash_cmd dep ver"
ECMD="dependence-version.sh mysqltest versoin"
HELP="❗ 依赖组件 \n格式: $FCMD\n例子: $ECMD"

# dep
if [ -z "${dep}" ]; then
	echo -e "${HELP}"
	exit
else
	echo -e "依赖组件 ${dep}"
fi

# ver
if [ -z "${ver}" ]; then
	echo -e "${HELP}"
	exit
else
	echo "依赖版本 ${ver}" # TODO 自动搜索 tests/*
fi

# 当前目录路径
WD="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "--> work dir:${WD}"

# 当前项目路径 pro
PD="$(cd "${WD}/../.." && pwd)"
echo "--> pro dir:${PD}"

# 目标目录
dirs=("${PD}/.github/workflows" "${PD}/eApi/internal/dao" "${PD}/eTest/internal/dao")
dirs="${dirs[@]}"
echo "--> dirs:${dirs}"

files="$(grep -rl "${dep}" ${dirs})"
files="${files[@]}"
echo "--> files:${files}"

sed -i "s/"${dep}:".*$/"${dep}"\:"${ver}"/g" ${files}
