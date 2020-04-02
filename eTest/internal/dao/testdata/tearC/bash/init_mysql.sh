#!/bin/bash
set -x
# set -e

set +x
echo "-----------------------init_mysql-----------------------------------"
set -x

uname -a

# mysql -uroot -proot -e "use mysql;update user set host = '%' where user ='root'"

# mysql -uroot -proot -e "CREATE USER 'utest'@'%' IDENTIFIED BY 'pwtest'"
# mysql -uroot -proot -e "grant all privileges on *.* to 'utest'@'%' with grant option"
# mysql -uroot -proot -e "flush privileges"
# mysql -uroot -proot -e "SELECT user,authentication_string,host FROM mysql.user"

# mysql -uroot -proot -e "create database dbtest"
# mysql -uroot -proot dbtest < sql/dbtest.sql
# mysql -uroot -proot -e "use dbtest;show tables"

# mysql -e "use mysql;update user set host = '%' where user ='root'"

mysql -uroot -e "CREATE USER 'utest'@'%' IDENTIFIED BY 'pwtest'"
mysql -e "grant all privileges on *.* to 'utest'@'%' with grant option"
mysql -e "flush privileges"
mysql -e "SELECT user,authentication_string,host FROM mysql.user"

mysql -e "create database dbtest"
mysql dbtest < sql/dbtest.sql
mysql -e "use dbtest;show tables"