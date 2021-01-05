#!/bin/bash
# set -x
set -e

echo -e "==> start check doc ..."

cd "$1"
# 处理改动的文件
for f in $2; do
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
