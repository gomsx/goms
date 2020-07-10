-- 创建用户
use mysql;
CREATE USER 'utest'@'%' IDENTIFIED BY 'pwtest';
grant all privileges on *.* to 'utest'@'%' IDENTIFIED BY 'pwtest' with grant option;
flush privileges;
SELECT user,authentication_string,host FROM mysql.user;

