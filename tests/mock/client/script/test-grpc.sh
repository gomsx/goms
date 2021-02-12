#!/bin/bash
set -x
set -e

# 当前目录路径
pwdx="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "==> pwdx:$pwdx"

while true; do
	bash $pwdx/test-api-grpc.sh "$1" "$2" "$3" "$4" "$5"
done
