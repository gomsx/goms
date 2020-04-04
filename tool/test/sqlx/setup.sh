#!/bin/bash
set -e

# echo `service mysql status`

echo ' ===> 1.启动mysql...'
service mysql start
sleep 3

echo `service mysql status`

echo ' ===> 2.创建库和表...'
mysql < /mysql/sqlx/test_db/create.sql

mysql < /mysql/sqlx/test_db/ping_table.sql
mysql < /mysql/sqlx/test_db/user_table.sql

mysql < /mysql/sqlx/test_db/ping_tablex.sql
mysql < /mysql/sqlx/test_db/user_tablex.sql

mysql < /mysql/sqlx/test_db/show_table.sql
echo '导入完毕...'

# sleep 3
# echo `service mysql status`

echo ' ===> 3.创建用户...'
mysql < /mysql/sqlx/user/create.sql

#sleep 3
echo `service mysql status`
echo ' ===> 4.初始化完成...'

tail -f /dev/null