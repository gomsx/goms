#!/bin/bash
set -x
set -e

# 工作目录路径
WD="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "--> work dir: ${WD}"

# 项目目录路径
PD="$(cd "$WD/.." && pwd)"
echo "--> pro dir: ${PD}"

# 子项目
SUBS=("eApi" "eTest" "eRedis" "eMysql" "eConf" "eGrpc" "eHttp" "eFlag" "eYaml" "eModule" "eDocker")

# work
for ((i = 0; i < ${#SUBS[*]}; i++)); do
	cd ${PD}/${SUBS[i]}/build && make $1 $2 $3 $4
done
