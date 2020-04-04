use mysql;
CREATE USER 'utest'@'%' IDENTIFIED BY 'pwtest';
grant all privileges on *.* to 'utest'@'%' with grant option;
flush privileges;
SELECT user,authentication_string,host FROM mysql.user;

create database dbtest;
dbtest < /sql/dbtest.sql;
use dbtest;show tables;