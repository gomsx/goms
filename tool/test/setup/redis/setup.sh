#!/bin/bash
set -x

echo "==> 1,start setup.sh"
PWD=$(cd "$(dirname "$0")";pwd)
echo "==> $PWD"

echo "==> 2,make redis.conf"
FILE=$HOME/.redis.conf
echo "requirepass pwtest" >> $FILE

echo "==> 3,start redis-server"
redis-server $FILE

