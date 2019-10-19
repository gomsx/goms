

```
go run .
go run . -confpath="../eConf/configs"
go run . -confpath="../../eConf/configs"


curl localhost:8080/ping

grpcurl -plaintext localhost:7070 list

grpcurl -plaintext -d '{"Message": "xxx"}'  localhost:7070 api.Test/Ping 

```


