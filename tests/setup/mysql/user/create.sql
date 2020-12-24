-- 创建用户
USE mysql;
CREATE USER 'utest'@'%' IDENTIFIED BY 'pwtest';
GRANT ALL PRIVILEGES ON *.* TO 'utest'@'%';
FLUSH PRIVILEGES;
SELECT user,authentication_string,host FROM mysql.user;

