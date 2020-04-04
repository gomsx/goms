#!/bin/bash
set -x
# set -e

set +x
echo "================= up_mysql_docker ======================="
set -x

## .sh所在目录
PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD

#run docker
docker --version

# docker run --name mysqltest -e MYSQL_ROOT_PASSWORD="pwroot" -e MYSQL_USER="utest" -e MYSQL_PASSWORD="pwtest" -p 31006:3306 -d mysql:5.7 
# -e MYSQL_DATABASE="xxx" -e MYSQL_USER="utest" -e MYSQL_PASSWORD="pwtest" 三个环境变量没效果
# docker run --name mysqltest -e MYSQL_ROOT_PASSWORD="root" -p 31006:3306 -d mysql:5.7 

DIR=$HOME
echo $DIR
sudo rm -rf $DIR/mysql

# 1) ok
# docker run --name mysqltest -p 13306:3306 -v/data/mysql/datadir:/var/lib/mysql -v /data/mysql/conf.d:/etc/mysql/conf.d -e MYSQL_ROOT_PASSWORD=root -d mysql:5.7

# 2) ok, 宿主机和容器的目录都不用提前手动创建
# docker run --name mysqltest -p 13306:3306 -v $DIR/mysql/data:/var/lib/mysql -v $DIR/mysql/conf:/etc/mysql/conf.d -e MYSQL_ROOT_PASSWORD=root -d mysql:5.7

# 3) ok, 分行
# docker run --name mysqltest \
#   -p 13306:3306 \
#   -v $DIR/mysql/data:/var/lib/mysql \
#   -v $DIR/mysql/conf:/etc/mysql/conf.d \
#   -eMYSQL_ROOT_PASSWORD=root \
#   -d mysql:5.7

#  4) ok
docker run --name mysqltest \
  -p 23306:3306 \
  -v $DIR/mysql/data:/var/lib/mysql \
  -v $DIR/mysql/conf:/etc/mysql/conf.d \
  -e MYSQL_ALLOW_EMPTY_PASSWORD  \
  -e MYSQL_ROOT_PASSWORD=root \
  -d mysql:5.7
echo " ==> mysqltest running"
# 观测点
# docker exec -it mysqltest /bin/bash

#cp file
docker cp $PWD/.my.cnf mysqltest:/root/
docker cp $PWD/sql mysqltest:/sql
docker cp $PWD/bash/init_mysql.sh mysqltest:/init_mysql.sh

docker exec -it mysqltest /bin/bash -c "chmod 644 /root/.my.cnf"
docker exec -it mysqltest /bin/bash -c "chmod a+x /init_mysql.sh"

echo " ==> file cp to mysqltest"
# 观测点
# docker exec -it mysqltest /bin/bash

docker restart mysqltest
sleep 3

# 观测点
# docker exec -it mysqltest /bin/bash

#run sh in docker 
COUNTER=0
while [ $COUNTER -lt 50 ]
do
    docker exec -it mysqltest /bin/bash -c "/init_mysql.sh"
    echo " ===> $COUNTER"
    let COUNTER+=1
done

echo " ==> run init_mysql.sh in docker"
# docker restart mysqltest

#ps docker
docker ps | grep mysqltest
