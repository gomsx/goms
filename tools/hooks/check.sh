#!/bin/bash
# set -x
# set -e

# echo -e "==> start check ..."
echo -e "\033[34m==> start check ...\033[0m"

# 解析参数
set="$1"

# 工作目录路径
WD="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# 项目目录路径
PD="$(cd "$WD/../.." && pwd)"

# 工具目录路径
TD=${PD}/tools/hooks

# TD_NAME 当前目录,工具集，不格式化
TD_NAME=$(basename "${TD}")

cd ${PD}

# 为项目中的 bash 文加上运行权限
find -name "*.sh" | xargs chmod +x

# cmd 获取改动的文件
files=
function change_files() {
	work_files="$(git diff --name-status --no-renames | grep -v "^D" | awk '{ print $2; }')"
	stage_files="$(git diff --name-status --cached --no-renames | grep -v "^D" | awk '{ print $2; }')"
	files="${work_files} ${stage_files}"
	# echo "$files"
}

# cmd 获取全部的文件，排除.git 和 工具目录
function all_files() {
	files="$(find -name "*" -type f | grep -v .git | grep -v ${TD_NAME})"
	# echo "$files"
}

# 文档排版检查
if [ "${set}" == "all" ]; then
	all_files
else
	change_files
fi
echo "--> files: $files"

# 文档排版检查
${TD}/check-doc.sh "${files}"

# go 代码静态检查
files_go=$(echo "$files" | grep .go)
echo "--> files go: ${files_go}"
if [ -n "${files_go}" ]; then
	${TD}/check-code-go.sh "${files_go}"
fi

# bash 代码静态检查
files_bash=$(echo "$files" | grep .sh)
echo "--> files bash: ${files_bash}"
if [ -n "${files_bash}" ]; then
	${TD}/check-code-bash.sh "${files_bash}"
fi

# echo -e "==< end check"
echo -e "\033[34m==< end check\033[0m"
