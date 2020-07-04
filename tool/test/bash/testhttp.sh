#!/bin/bash
set -x

PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD

while true
do
	# bash $PWD/test_http.sh $1 $2 $3			# 如果 $1 为 ""(即空),那么 test_http.sh 无法获取到
	bash $PWD/test_http.sh "$1" "$2" "$3"		# 如果 $1 为 ""(即空),那么 test_http.sh 可以获取到 
	sleep 1
done

