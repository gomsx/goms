#!/bin/bash
# set -x
set -e

# echo -e "==> start check ..."
echo -e "\033[34m==> start check ...\033[0m"

# 当前 bash 所在目录路径 PWD
PWD=$(cd "$(dirname "$0")";pwd)

# 当前项目路径 PRO
PRO=$PWD/../..
PRO=$(cd $PRO;pwd)

# 为项目中的 bash 文加上运行权限
find $PRO -name "*.sh" | xargs chmod +x 

# 文档排版检查
[ "$1" ] || $PWD/ck-doc-deta.sh
[ "$1" = "all" ] && $PWD/ck-doc-all.sh

# 代码静态检查
$PWD/ck-code-go.sh

# echo -e "==< end check"
echo -e "\033[34m==< end check\033[0m"

