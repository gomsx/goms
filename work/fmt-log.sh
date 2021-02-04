#!/bin/bash
set -x
set -e

# 工程 root
cd ..

# 搜索 go 文件
dst="*.go"
files="$(find -name "$dst")"
echo "--> files: $files"

# 处理文件
for file in $files; do
	# 运行 2 次
	sed -i "/log\..*)\.$/{ N; s/\n// }" "$file"
	sed -i "/log\..*)\.$/{ N; s/\n// }" "$file"
done

go fmt ./...
