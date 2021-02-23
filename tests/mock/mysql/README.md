## build images

```
make docker
```

## run container

```
docker run --name msyqlx -p 13306:3306 -d mysqltest
```

## login mysql (port:13306, user:utest, password:pwtest)

```
mysql -h127.0.0.1 -P13306 -uutest -ppwtest
mysql -h172.17.0.1 -P13306 -uutest -ppwtest
```
## login container
```
docker exec -it msyqlx /bin/bash
```

## 调试 mysqltest

```
## run
docker run --name msyqlx -p 13306:3306 -d mysqltest

## set
mysql -h127.0.0.1 -P13306 -uroot < test_db/ping_table.sql
mysql -h127.0.0.1 -P13306 -uroot < test_db/user_table.sql
mysql -h127.0.0.1 -P13306 -uroot -Bse 'show databases'

## login mysql
mysql -h127.0.0.1 -P13306 -uroot

## exec login container
docker exec -it msyqlx /bin/bash
mysql -Bse 'show databases'
```

## 调试 mysql8

```
## run
docker run --name mysqlx8 -p 13306:3306 \
-e MYSQL_ROOT_PASSWORD=123456 \
-e MYSQL_USER=sfw \
-e MYSQL_PASSWORD=123456 \
-e MYSQL_DATABASE=test_db \
-d mysql:8.0

## set
mysql -h127.0.0.1 -P13306 -usfw -p < test_db/user_table.sql

## login msyql
mysql -h127.0.0.1 -P13306 -usfw -p

## exec login container
docker exec -it msyqlx8 /bin/bash
```

>
https://dev.mysql.com/doc/refman/8.0/en/docker-mysql-more-topics.html#docker-persisting-data-configuration  
https://jingyan.baidu.com/article/0bc808fcbc4b155bd485b9cb.html  
