#!/bin/bash
set -x
# set -e

set +x
echo "-----------------------init_mysql-----------------------------------"
set -x

uname -a
service mysql restart 
service mysql status

# mysql -uroot -ppwroot -e "grant all privileges on *.* to 'utest'@'%' with grant option"
# mysql -uroot -ppwroot -e "flush privileges"
# mysql -uutest -ppwtest -e "create database dbtest"
# mysql -uutest -ppwtest dbtest < sql/dbtest.sql

mysql -e "grant all privileges on *.* to 'utest'@'%' with grant option"
mysql -e "flush privileges"

mysql -e "create database dbtest"
mysql dbtest < sql/dbtest.sql
mysql -e "show databases"