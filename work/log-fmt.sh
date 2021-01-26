#!/bin/bash
set -x
# set -e

# 当前目录路径
pwdx=$(
	cd "$(dirname "$0")"
	pwd
)

# 当前项目路径 pro
pro=$(
	cd "$pwdx/.."
	pwd
)
echo "==> pro_dir:$pro"

# exit

# 搜索 go 文件
dst="*.go"
cmd="find \"$pro\" -name \"$dst\""
files=$(
	cd "$pro"
	eval "$cmd"
)

echo "--> files: $files"
file_set=("$files")
count=${#file_set[*]}
echo "--> count: $count"

# exit

# 处理文件
for file in $files; do

	# 2 替换 log.Ctx(c).xxx(). 为 log.xxxf(
	old="log.Ctx(c).Info()."
	old1="log.Ctx(c).Error()."
	old2="log.Ctx(c).Warn()."
	old3="log.Ctx(c).Debug()."

	new="log.Ctx(c).Info("
	new1="log.Ctx(c).Error("
	new2="log.Ctx(c).Warn("
	new3="log.Ctx(c).Debug("

	sed -i "s/$old/$new/g" "$file"
	sed -i "s/$old1/$new1/g" "$file"
	sed -i "s/$old2/$new2/g" "$file"
	sed -i "s/$old3/$new3/g" "$file"

	# 3 替换 Msgf 为 log.Printf
	old="\tMsgf(\""
	old1="\tMsg(\""
	new=").Msgf(\""
	new1=").Msg(\""
	sed -i "s/$old/$new/g" "$file"
	sed -i "s/$old1/$new1/g" "$file"

	# 3 替换 Msgf 为 log.Printf
	# old="().)"
	# new="()"
	# sed -i "s/$old/$new/g" "$file"
done
