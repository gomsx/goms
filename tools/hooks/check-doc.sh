#!/bin/bash
# set -x
# set -e

echo -e "==> start check doc ..."

files="$1"

# 处理改动的文件
for file in $files; do
	# 行尾，删除行尾空格
	[ "${file##*.}" != "md" ] && (sed -i 's/ *$//' "$file")
	
	# 行首，替换特殊字符(空格、tab等)成换行符
	sed -i 's/^\s*$/\n/g' "$file"

	# 尾行 1，任意不是'\n'(正文),则追加'\n'(换行)
	# ecoh -n 不附加换行符
	# 字符串判断(=~ 正则表达式)(== 通配符)
	[[ "$(tail -1 "$file")" =~ .+[^\n]$ ]] && (echo -e -n "\n" >>  "$file")
	
	# 尾行 2，是'\n',则删除
	while [ "$(wc -l < "$file")" -gt 0 ] && [ "$(tail -1 "$file")" == "" ]; do
		mv "$file" "$file".bak
		tac "$file".bak | sed '1d' | tac > "$file"
		rm -f "$file".bak
	done

	# 首行，删除为空的首行
	sed -i '/./,$!d' "$file"

	# 合并多个空行
	sed -i '/^$/{N;/^\n*$/D}' "$file"	
done

echo -e "==< end check doc"
