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
	cd "$pwdx/../.."
	pwd
)

# 搜索包含 pb.go 的文件
dst=pb.go
find_dir="find \"$pro\" -name \"$dst\""
files=$(eval "$find_dir")

file_set=("$files")
count=${#file_set[*]}
echo "--> count: $count"

# 处理文件
for f in $files; do
    # 1
    path=$(dirname "$f")

    # 2
    old="option go_package = \""
    new="option go_package = \"..\/"
    gf=$(grep -rl "$old" "$path")
    sed -i "s/$old/$new/g" "$gf"

    # 3
    cd "$path";go generate

    # 4
    old="_ \"google\/api\""
    new="_ \"google.golang.org\/genproto\/googleapis\/api\/annotations\""
    gf=$(grep -rl "$old" "$path"/..)
    sed -i "s/$old/$new/g" "$gf"
done