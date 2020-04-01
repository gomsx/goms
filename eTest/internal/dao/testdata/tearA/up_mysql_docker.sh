#!/bin/bash
set -x
# set -e

## .sh所在目录
DIR=$(cd "$(dirname "$0")";pwd)
echo $DIR

#run docker
docker --version
docker pull mysql:5.7
docker run --name mysqltest -e MYSQL_ROOT_PASSWORD="pwroot" -p 31006:3306 -d mysql:5.7 

#cp file
docker cp $DIR/.my.cnf mysqltest:/root/
docker cp $DIR/sql mysqltest:/sql
docker cp $DIR/bash/init_mysql.sh mysqltest:/init_mysql.sh
docker exec -it mysqltest /bin/bash -c "chmod 644 /root/.my.cnf"
docker restart mysqltest
sleep 3
#run sh in docker 
docker exec -it mysqltest /bin/bash -c "chmod +x init_mysql.sh && /init_mysql.sh"
docker restart mysqltest

#ps docker
docker ps | grep mysqltest