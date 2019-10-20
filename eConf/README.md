


```
go run .
go run . -confpath="../eConf/configs"
go run . -confpath="../../eConf/configs"


curl  localhost:8080/test/ping

grpcurl -plaintext localhost:50051 list

grpcurl -plaintext -d '{"Message": "xxx"}'  localhost:50051 api.Test/Ping 

```