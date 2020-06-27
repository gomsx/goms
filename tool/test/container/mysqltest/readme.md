

## build images
```
make docker
```

## run container
```
docker run -p 13306:3306 -d mysqltest
```

## login (port:13306, user:utest, password:pwtest)
```
mysql -h127.0.0.1 -P13306 -uutest -ppwtest
mysql -h172.17.0.1 -P13306 -uutest -ppwtest
```
