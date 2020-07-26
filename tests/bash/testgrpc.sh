#!/bin/bash
set -x

PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD

echo "$1"
echo "$2"

while true
do
	bash $PWD/test_grpc.sh "$1" "$2" "$3" "$4"
done

