#!/bin/bash
set -x
# set -e

set +x
echo "-----------------------init_mysql-----------------------------------"
set -x

uname -a

mysql -e "CREATE USER 'utest'@'%' IDENTIFIED BY 'pwtest'"
mysql -e "grant all privileges on *.* to 'utest'@'%' with grant option"
mysql -e "flush privileges"
mysql -e "SELECT user,authentication_string,host FROM mysql.user"

mysql -e "create database dbtest"
mysql dbtest < sql/dbtest.sql
mysql -e "use dbtest;show tables"