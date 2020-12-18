## build images
```
make docker
```

## run container
```
docker run -p 16379:6379 -d goms-redistest
```

## login (port:16379, password:pwtest)
```
redis-cli -p 16379 -a pwtest
```

