# dlv

```
# run app
eApi

# print pid
ps aux | grep eApi | awk 'NR==1 { print $2 }'

# attach app
dlv attach xxx

# list code
l

# breakpoint
b /home/fws/github.com/vuca/goms/eApi/internal/service/ping.go:13

# breakpoint print
bp

# continue
c

# next line
n

# step
s

# step out
so

# local var print
locals msg
```

```
grpcurl  -max-time 9999999 -plaintext localhost:50051 service.goms.v1.User/Ping
```
