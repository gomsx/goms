#!/bin/bash
set -x

PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD

while true
do
	bash $PWD/test_http.sh $1 $2
	sleep 1
done
