

```
curl localhost:8080/ping

grpcurl -plaintext localhost:7070 list

grpcurl -plaintext -d '{"Message": "gopher"}'  localhost:7070 pb.Egrpc/Ping

```


