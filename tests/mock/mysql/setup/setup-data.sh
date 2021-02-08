#!/bin/bash
set -x
set -e

echo ' ==> 创建库和表...'
mysql -uroot <test_db/create.sql

mysql -uroot <test_db/ping_table.sql
mysql -uroot <test_db/user_table.sql

mysql -uroot <test_db/show_table.sql

echo ' ==> 创建用户...'
mysql -uroot <user/create.sql
