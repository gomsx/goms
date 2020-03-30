#!/bin/bash
set -x
# set -e

## .sh所在目录
DIR=$(cd "$(dirname "$0")";pwd)
echo $DIR

#run docker
docker --version
mysql --version
docker pull mysql:5.7
docker run --name mysqltest -e MYSQL_ROOT_PASSWORD="pwroot" -e MYSQL_USER="utest" -e MYSQL_PASSWORD="pwtest" -p 31006:3306 -d mysql:5.7 

#delay
sleep 5 
docker ps | grep mysqltest

#cp file

docker cp $DIR/.my.cnf mysqltest:/root/
docker cp $DIR/sql mysqltest:/sql
docker cp $DIR/bash/init_mysql.sh mysqltest:/init_mysql.sh
docker exec -it mysqltest /bin/bash -c "chmod 644 /root/.my.cnf"

#run sh in docker 
docker exec -it mysqltest /bin/bash -c "chmod +x init_mysql.sh && /init_mysql.sh"

#ps docker
docker ps | grep mysqltest