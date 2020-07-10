#!/bin/bash
set -x
[ $1 ] && SERVICE="service.goms.$1" ||SERVICE="service.goms" 
[ $2 ] && HOST=$2 || HOST=localhost
[ $3 ] && PORT=$3 || PORT=50051

ADDR="$HOST:$PORT"

# ping
# Ping
grpcurl -plaintext $ADDR $SERVICE.User/Ping 

# Ping
grpcurl -plaintext -d '{"message": "xxx"}' $ADDR $SERVICE.User/Ping 

# user
# CreateUser
res=$(grpcurl -plaintext -d '{"name": "xxx","sex":"0"}' $ADDR $SERVICE.User/CreateUser)
res=$(echo $res | awk 'NR==1{ print $3 }' | tr -d "\"")
uid=$res;
name=name${uid:1:6};echo $name

# UpdateUser 
data='{"uid":"=uid","name":"=name","sex":"1"}'
data=$(echo $data | sed s/=uid/$uid/ |sed s/=name/$name/)
grpcurl -plaintext -d $data $ADDR $SERVICE.User/UpdateUser

# ReadUser
data='{"uid":"=uid"}'
data=$(echo $data | sed s/=uid/$uid/)
grpcurl -plaintext -d $data $ADDR $SERVICE.User/ReadUser

# DeleteUser
grpcurl -plaintext -d $data $ADDR $SERVICE.User/DeleteUser

