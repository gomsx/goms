#!/bin/bash
set -x
set -e

# 当前目录 
WD=$(dirname "$0")
WD=$(cd "$WD"; pwd)

# 项目路径 PD
PD="$(cd ${WD}/.. && pwd)"
echo "==> PD:${PD}"

# 搜索 go 文件
dst="*.go"
files="$(find ${PD} -name ${dst})"
echo "--> files: ${files}"

# 处理文件
for file in $files; do

	# delete RequestId
	old="setRequestId"
	sed -i "/$old/d" ${file}

	old1="ctxq"
	old2="ctxu"
	new="ctxb"
	sed -i "s/$old1/$new/g" ${file}
	sed -i "s/$old2/$new/g" ${file}

	# delete UserId
	old="ms.CarryCtxUserId"
	sed -i "/$old/d" ${file}

	old="ms.GetCtxVal"
	sed -i "/$old/d" ${file}

	# c -> ctx
	old="(c context"
	new="(ctx context"
	sed -i "s/$old/$new/g" ${file}

	old="(c,"
	new="(ctx,"
	sed -i "s/$old/$new/g" ${file}

	old="(c)"
	new="(ctx)"
	sed -i "s/$old/$new/g" ${file}

	# patch
	old="Notify(ctx,"
	new="Notify(c,"
	sed -i "s/$old/$new/g" ${file}
done
