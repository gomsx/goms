#!/bin/bash
# set -x
set -e

echo -e "==> start check doc ..."

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

# pwdx_name 当前目录,工具集，不格式化
pwdx_name=$(basename "$pwdx")

# cmd 获取改动的文件
cmd_deta="git status -s | awk '{ print \$2; }' | grep -v /$pwdx_name" # $2 要做字符串处理，即 \$2

# cmd 获取要格式化的文件，排除.git 和 工具目录
cmd_all="find $pro -name \"*\" -type f | grep -v /.git | grep -v /$pwdx_name"

# 文档排版检查
[ "$1" ] || cmd=$cmd_deta
[ "$1" = "all" ] && cmd=$cmd_all

# files 要格式化的文集合
files=$(eval "$cmd")
echo "--> files: $files"

# 处理改动的文件
for f in $files; do
	# 匹配空格、tab等特殊字符,替换成换行符
	sed -i 's/^\s*$/\n/g' "$f"
	# 尾行部插入空行
	sed -i '$a\\n' "$f"
	# 合并多个空行
	sed -i '/^$/{N;/^\n*$/D}' "$f"
	# 删除为空的首行
	sed -i '/./,$!d' "$f"
done

echo -e "==< end check doc"
