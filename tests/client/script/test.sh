#!/bin/bash
set -x
set -e

# 当前目录路径
pwdx="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "==> pwdx:$pwdx"

time=$1
for(( i=0; i<$time; i++));
do
	bash $pwdx/test-api-http.sh "$2" "$3" "$4" "$5" "$6"
	bash $pwdx/test-api-grpc.sh "$2" "$3" "$4" "$5" "$7"
done
