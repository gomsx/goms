#!/bin/bash
# set -x
set -e

echo -e "==> start xcheck ..."

# 目录路径
WD="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# 项目路径
PD="$(cd "${WD}/../.." && pwd)"

# 工具目录路径
TD=${PD}/tools/hooks

# TD_NAME 当前目录,工具集，不格式化
TD_NAME=$(basename "${TD}")

# 用 new 替换 old
old="aivuca"
new="fuwensun"

# cmd 搜索包含 old 的文件
cmd="grep ${old} -rl --exclude-dir={.git,${TD_NAME}}"
cmde="grep ${old} -rl"

# files 包含 old 的文件集合
echo "替换前："
files=$(cd ${PD} && eval "${cmd}")
echo "--> files: ${files}"

# 执行替换
echo "替换："
cd ${PD} && sed -i "s/${old}/${new}/g" ${files} # $files ==> f1 f2 ... # "$files" ==> 'f1 f2'

echo "替换后："
cd ${PD} && eval "${cmde}"

echo -e "==< end xcheck"
