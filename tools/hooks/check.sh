#!/bin/bash
# set -x
# set -e

# echo -e "==> start check ..."
echo -e "\033[34m==> start check ...\033[0m"

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

# 工具目录 toolx
toolx=$pro/tools/hooks

# 为项目中的 bash 文加上运行权限
find "$pro" -name "*.sh" | xargs chmod +x

# 文档排版检查
"$toolx"/check-doc.sh "$1"

# 代码静态检查
"$toolx"/check-code-go.sh

# echo -e "==< end check"
echo -e "\033[34m==< end check\033[0m"
