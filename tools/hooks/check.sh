#!/bin/bash
# set -x
set -e

# echo -e "==> start check ..."
echo -e "\033[34m==> start check ...\033[0m"

# 当前 bash 所在目录路径 pwdx
pwdx=$(
	cd "$(dirname "$0")"
	pwd
)

# 当前项目路径 pro
pro=$pwdx/../..
pro=$(
	cd "$pro"
	pwd
)

# 为项目中的 bash 文加上运行权限
find "$pro" -name "*.sh" | xargs chmod +x

# 文档排版检查
"$pwdx"/ck-doc.sh "$1"

# 代码静态检查
"$pwdx"/ck-code-go.sh

# echo -e "==< end check"
echo -e "\033[34m==< end check\033[0m"
