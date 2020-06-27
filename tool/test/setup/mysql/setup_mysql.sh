#!/bin/bash
set -e

PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD


echo ' ==> 创建库和表...'
mysql < $PWD/test_db/create.sql

mysql < $PWD/test_db/ping_table.sql
mysql < $PWD/test_db/user_table.sql

mysql < $PWD/test_db/show_table.sql

echo ' ==> 创建用户...'
mysql < $PWD/user/create.sql
