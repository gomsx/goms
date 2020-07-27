#!/bin/bash
set -x

[ $1 ] && US=$1 || US=100
[ $2 ] && SERVICE="service.goms.$2" ||SERVICE="service.goms" 
[ $3 ] && HOST=$3 || HOST=localhost
[ $4 ] && PORT=$4 || PORT=50051

ADDR="$HOST:$PORT"

# usleep : 默认以微秒。  
# 1s = 1000ms = 1000000us
function delay(){
    # sleep 1
    usleep $US
    echo "==> delay $US us"
}

# ping
# Ping
grpcurl -plaintext $ADDR $SERVICE.User/Ping 

# Ping
grpcurl -plaintext -d '{"message":"xxx"}' $ADDR $SERVICE.User/Ping 

# user
# CreateUser
data='{"name":"xxx","sex":"1"}'
CMD="grpcurl -plaintext -d \$data \$ADDR \$SERVICE.User/CreateUser"

res=$(eval $CMD)

CMD="echo $res | awk 'NR==1{ print \$3 }' | tr -d \"\"\""
res=$(eval $CMD)
uid=$res
name=name${uid:1:6}

delay

# UpdateUser 
data='{"uid":"=uid","name":"=name","sex":"1"}'
data=$(echo $data | sed s/=uid/$uid/ |sed s/=name/$name/)
grpcurl -plaintext -d $data $ADDR $SERVICE.User/UpdateUser

delay

# ReadUser
data='{"uid":"=uid"}'
data=$(echo $data | sed s/=uid/$uid/)
grpcurl -plaintext -d $data $ADDR $SERVICE.User/ReadUser

delay

# DeleteUser
grpcurl -plaintext -d $data $ADDR $SERVICE.User/DeleteUser

delay

