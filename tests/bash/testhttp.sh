#!/bin/bash
set -x
set -e

pwdx=$(
	cd "$(dirname "$0")"
	pwd
)

while true; do
	# 当 $1 为 "" 时，此处 "$1" 是 "" ,而 $1 是空白（noting）
	bash $pwdx/test_http.sh "$1" "$2" "$3" "$4"
done
