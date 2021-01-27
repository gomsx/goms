#!/bin/bash
set -x
# set -e

# 当前目录路径
pwdx="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "==> pwdx:$pwdx"

# 当前项目路径 pro
pro="$(cd "$pwdx/.." && pwd)"
echo "==> pro:$pro"

# 搜索 go 文件
dst="*.go"
files="$(find "$pro" -name "$dst")"
echo "--> files: $files"

file_set=("$files")
count=${#file_set[*]}
echo "--> count: $count"

# 处理文件
for file in $files; do

	# 1 替换 log.Ctx(c). 为 log.
	old="log.Ctx(c)."
	new="log."
	sed -i "s/$old/$new/g" "$file"

	# 2 替换 log.xxxf(). 为 log.xxxf(
	old="log.Info()..*Msg"
	old1="log.Error()..*Msg"
	old2="log.Warn()..*Msg"
	old3="log.Debug()..*Msg"
	old4="log.Fatal()..*Msg"

	new="log.Info"
	new1="log.Error"
	new2="log.Warn"
	new3="log.Debug"
	new4="log.Fatal"

	sed -i "s/$old/$new/g" "$file"
	sed -i "s/$old1/$new1/g" "$file"
	sed -i "s/$old2/$new2/g" "$file"
	sed -i "s/$old3/$new3/g" "$file"
	sed -i "s/$old4/$new4/g" "$file"

	# 3 num, 为 _,
	old="num, err :="
	new="_, err ="
	sed -i "s/$old/$new/g" "$file"

	# 4 替换 import pkg
	# no
	# old='"github.com\rs\zerolog\log"'
	# new='log "github.com\sirupsen\logrus"'
	# ok
	old="\"github.com\/rs\/zerolog\/log\""
	new="log \"github.com\/sirupsen\/logrus\""
	sed -i "s/$old/$new/g" "$file"
done
