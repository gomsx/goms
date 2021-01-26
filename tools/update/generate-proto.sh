#!/bin/bash
set -x
# set -e

# 当前目录路径
pwdx="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "==> pwdx:$pwdx"

# 当前项目路径 pro
pro="$(cd "$pwdx/../.." && pwd)"
echo "==> pro:$pro"

# 搜索包含 pb.go 的文件
dst=pb.go
find_dir="find \"$pro\" -name \"$dst\""
files=$(eval "$find_dir")

file_set=("$files")
count=${#file_set[*]}
echo "--> count: $count"

# 处理文件
for file in $files; do
    
    # 1 pb.go 所在的目录
    path=$(dirname "$file")

    # 2 替换 go pkg 路径
    # old="option go_package = \""
    # new="option go_package = \"..\/"
    # genfile=$(grep -rl "$old" "$path")
    # sed -i "s/$old/$new/g" "$genfile"

    # 3 编译 proto 文件
    cd $path;go generate

    # 4 替换 import pkg
    old="_ \"google\/api\""
    new="_ \"google.golang.org\/genproto\/googleapis\/api\/annotations\""
    gofile=$(grep -rl "$old" "$path"/..)
    sed -i "s/$old/$new/g" "$gofile"
done