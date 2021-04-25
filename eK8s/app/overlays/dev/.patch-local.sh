#!/bin/bash
set -x
set -e

FS=$(grep -rl "/*" | grep -v "patch-local.sh")

sed -i "s/\/var\/lib\/mysqlx/\/var\/lib\/mysql-x/g" $FS
sed -i "s/\/var\/lib\/redisx/\/var\/lib\/redis-x/g" $FS

sed -i "s/name\: pv-local-mysql/name\: fuwensun-pv-local-mysql/g" $FS
sed -i "s/name\: pv-local-redis/name\: fuwensun-pv-local-redis/g" $FS
