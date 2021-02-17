#!/bin/bash
set -e
# set -x

cm="$(git log -1 --oneline)"
cm=${cm,,} # 转小写

keys=("client" "mysql" "redis")
keys="${keys[@]}"

for key in ${keys};do
    msg="update ${key}test"
    if [[ "${cm}" == *"${msg}"* ]];then
        echo "==> build: ${key}"
        (cd ../${key}/build && make docker push)
    else
        echo "==> pass: ${key}"
    fi
done
