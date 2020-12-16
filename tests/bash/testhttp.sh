#!/bin/bash
set -e
set -x

PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD

while true
do
	# 当 $1 为 "" 时，此处 "$1" 是 "" ,而 $1 是空白（noting）
	bash $PWD/test_http.sh "$1" "$2" "$3" "$4" 
done

