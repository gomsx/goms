#!/bin/bash
set -x
set -e

# 当前目录路径
pwdx=$(
	cd "$(dirname "$0")"
	pwd
)

echo ' ==> 创建库和表...'
mysql -uroot < "$pwdx"/test_db/create.sql

mysql -uroot < "$pwdx"/test_db/ping_table.sql
mysql -uroot < "$pwdx"/test_db/user_table.sql

mysql -uroot < "$pwdx"/test_db/show_table.sql

echo ' ==> 创建用户...'
mysql -uroot < "$pwdx"/user/create.sql
