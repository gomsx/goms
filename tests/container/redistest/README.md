## build images
```
make docke
```

## run container
```
docker run -p 16379:6379 -d redistest
```

## login (port:16379, password:pwtest)
```
redis-cli -p 16379 -a pwtest
```

