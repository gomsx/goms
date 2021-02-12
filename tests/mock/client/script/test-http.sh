#!/bin/bash
set -x
set -e

# 当前目录路径
pwdx="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "==> pwdx:$pwdx"

while true; do
	# 当 $1 为 "" 时，此处 "$1" 是 "" ,而 $1 是空白（noting）
	bash $pwdx/test-api-http.sh "$1" "$2" "$3" "$4" "$5"
done
