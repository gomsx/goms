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
	# 例子：log.Ctx(c).Info().Int64("rows", num).Msgf("db insert user: %v", *user)

	# 1、"Int64("rows", num)." 移动到行末尾，并注释掉
	# 例子：log.Ctx(c).Info().Msgf("db insert user: %v", *user)//Int64("rows", num).
	sub1="log\..*()\."
	sub2x="Int64(.*)."
	sub2y="Str(.*)."
	sub3="Msg.*)"
	# debug
	# sed -n "/$sub1/p" $file
	# sed -n "/$sub2x/p" $file
	# sed -n "/$sub2y/p" $file
	# sed -n "/$sub3/p" $file
	sed -i "s/\($sub1\)\($sub2x\)\($sub3\)/\1\3\/\/\2/" $file
	sed -i "s/\($sub1\)\($sub2y\)\($sub3\)/\1\3\/\/\2/" $file

	# 2、删除 "Ctx(x)."
	# 例子：log.Info().Msgf("db insert user: %v", *user)//Int64("rows", num).
	old1="Ctx()."
	old2="Ctx(c)."
	sed -i "s/$old1//" $file
	sed -i "s/$old2//" $file

	# 3、删除 "().Msg"
	# 例子：log.Infof("db insert user: %v", *user)//Int64("rows", num).
	old="()\.Msg"
	sed -i "s/$old//" $file

	# 4、替换 import pkg
	# github.com\rs\zerolog\log ==> github.com\sirupsen\logrus
	old="\"github\.com\/rs\/zerolog\/log\""
	new="log \"github\.com\/sirupsen\/logrus\""
	sed -i "s/$old/$new/g" $file
	# 修正 log log "github.com\sirupsen\logrus" 情况
	sed -i "s/log log/log/g" $file

	# log 细节的替换
	old="\.GlobalLevel"
	new="\.GetLevel"
	sed -i "s/$old/$new/g" $file
	old="\.SetGlobalLevel"
	new="\.SetLevel"
	sed -i "s/$old/$new/g" $file
	old="zerolog"
	new="log"
	sed -i "s/$old/$new/g" $file

	# 5、替换 RowsAffected() 返回值 num 为 _
	old="num, err :="
	new="_, err ="
	sed -i "/RowsAffected/{ s/$old/$new/ }" $file

	# 6、删除包含 log.Ctx 的行
	old="log\.Ctx("
	sed -i "/$old/d" $file
done
