#!/bin/bash
set -x
set -e

# 当前目录路径
WD="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "==> work dir: ${WD}"

while true; do
	bash ${WD}/test-api-grpc.sh "$1" "$2" "$3" "$4" "$5"
done
