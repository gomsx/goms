#!/bin/bash
set -x

PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD

while true
do
	bash $PWD/test_http.sh $1 $2
	bash $PWD/test_grpc.sh $1 $3
    sleep 1
done
