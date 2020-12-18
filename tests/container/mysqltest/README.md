## build images
```
make docker
```

## run container
```
docker run -p 13306:3306 -d goms-mysqltest
```

## login (port:13306, user:utest, password:pwtest)
```
mysql -h127.0.0.1 -P13306 -uutest -ppwtest
mysql -h172.17.0.1 -P13306 -uutest -ppwtest
```

## 调试 goms-mysqltest
```
## run
docker run -d --mydbx -p 13306:3306 goms-mysqltest:latest

## set
mysql -h127.0.0.1 -uroot < test_db/ping_table.sql
mysql -h127.0.0.1 -uroot < test_db/user_table.sql

## login
mysql -h127.0.0.1 -uroot

## exec
docker exec -it mydbx /bin/bash

mysql -Bse 'show databases'
```

## 调试 mysql:8
```
## run
sudo docker run  -d --name mydbxx -p 13306:3306 -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_USER=sfw -e MYSQL_PASSWORD=123456 -e MYSQL_DATABASE=test_db mysql:8.0

## set
mysql -h127.0.0.1 -P13306 -usfw -p < test_db/user_table.sql

## login
mysql -h127.0.0.1 -P13306 -usfw -p

## exec
docker exec -it mydbxx /bin/bash
```

>https://dev.mysql.com/doc/refman/8.0/en/docker-mysql-more-topics.html#docker-persisting-data-configuration  
https://jingyan.baidu.com/article/0bc808fcbc4b155bd485b9cb.html  