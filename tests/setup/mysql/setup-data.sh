#!/bin/bash
set -x
set -e

# 当前目录路径
pwdx=$(
	cd "$(dirname "$0")"
	pwd
)

echo ' ==> 创建库和表...'
mysql < "$pwdx"/test_db/create.sql

mysql < "$pwdx"/test_db/ping_table.sql
mysql < "$pwdx"/test_db/user_table.sql

mysql < "$pwdx"/test_db/show_table.sql

echo ' ==> 创建用户...'
mysql < "$pwdx"/user/create.sql

