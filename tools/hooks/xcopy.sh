#!/bin/bash
# set -x
set -e

echo -e "==> start xcopy ..."

# cp_dst 要复制的目标
cp_dst="$1"

HELP="❗ 错误，缺少目标参数"
FCMD="格式: bash_cmd copy_dst"
ECMD="例子: xcopy.sh xx/yy"
# 如果没有目标参数，打印提示并退出
if [ -z "${cp_dst}" ]; then
	echo -e "$HELP"
	echo -e "$FCMD"
	echo -e "$ECMD"
	exit
else
	echo "复制目标 ${cp_dst}"
fi

# 当前目录路径
WD="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# 当前项目路径 pro
PD="$(cd "${WD}/../.." && pwd)"

# 用 fromx/... 替换 tox/...
fromx="aivuca"
tox="fuwensun"

# dst 目标，src 源头
dst=${PD}/${cp_dst}
src=${dst/$tox/$fromx}

# 执行 copy
rm -rf "${dst}"
cp -r "${src}" "${dst}"

# 执行 xcheck
(cd ${PD}/tools/hooks && bash xcheck.sh)

echo -e "==< end xcopy"
