#!/bin/bash
set -x
set -e

# 当前目录路径
WD="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "--> work dir:${WD}"

# 当前项目路径
PD="$(cd "${WD}/../.." && pwd)"
echo "--> pro dir:${PD}"

# 替换 proto 文件中 go pkg 路径
# files="$(find "${PD}" -name "*.proto" | grep -v pkg)"
# echo "${files}"
# old="^option go_package = \""
# ins="\.\.\/"
# sed -i "s/${old}/${old}${ins}/g" ${files}
# exit

# 搜索包含 pb.go 的文件,目录
dst=pb.go
files="$(find "${PD}" -name "${dst}")"
echo "${files}"
dirs="$(echo "${files}" | xargs dirname)"
echo "${dirs}"

# 处理文件
for dir in ${dirs}; do
    # 1 generate
    (cd ${dir} && go generate)

    # 2 替换 import pkg
    old="_ \"google\/api\""
    new="_ \"google.golang.org\/genproto\/googleapis\/api\/annotations\""
    gofile=$(grep -rl "${old}" ${dir}/..)
    sed -i "s/${old}/${new}/g" "${gofile}"
done