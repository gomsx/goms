#!/bin/bash
set -x
set -e

# 项目 root
PROX="$(cd "../" && pwd)"

# 子项目
SUBS=("eApi" "eTest" "eRedis" "eMysql" "eConf" "eGrpc" "eHttp" "eFlag" "eYaml" "eModule" "eDocker")

# work
for ((i = 0; i < ${#SUBS[*]}; i++)); do
	cd ${PROX}/${SUBS[i]}/build && make compile docker push clean
done
