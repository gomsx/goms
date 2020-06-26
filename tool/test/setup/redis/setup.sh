#!/bin/bash

set -x

PWD=$(cd "$(dirname "$0")";pwd)

FILE=$PWD/redis.conf

service redis stop

redis-server $FILE

