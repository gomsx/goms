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

# toolx_name 当前目录,工具集，不格式化
toolx_name=$(basename "$toolx")

# 为项目中的 bash 文加上运行权限
find "$pro" -name "*.sh" | xargs chmod +x

# cmd 获取改动的文件
cmd_deta="git status -s | awk '{ print \$2; }' | grep -v /$toolx_name" # $2 要做字符串处理，即 \$2

# cmd 获取要格式化的文件，排除.git 和 工具目录
cmd_all="find $pro -name \"*\" -type f | grep -v /.git | grep -v /$toolx_name"

# 文档排版检查
[ "$1" ] || cmd=$cmd_deta
[ "$1" = "all" ] && cmd=$cmd_all

# files 要格式化的文集合
files=$(
	cd "$pro"
	eval "$cmd"
)
echo "--> files: $files"

# 文档排版检查
"$toolx"/check-doc.sh "$pro" "$files"

# 代码静态检查
files_go=$(echo "$files" | grep .go)
echo "--> files go: $files_go"
if [ -n "$files_go" ]; then
	"$toolx"/check-code-go.sh "$pro" "$files_go"
fi

# 代码静态检查
files_bash=$(echo "$files" | grep .sh)
echo "--> files bash: $files_bash"
if [ -n "$files_bash" ]; then
	"$toolx"/check-code-bash.sh "$pro" "$files_bash"
fi

# echo -e "==< end check"
echo -e "\033[34m==< end check\033[0m"
