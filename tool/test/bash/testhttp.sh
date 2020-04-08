#!/bin/bash
set -x

PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD

while true
do
	bash $PWD/test_http.sh
	# bash $PWD/test_grpc.sh
	sleep 1
done
