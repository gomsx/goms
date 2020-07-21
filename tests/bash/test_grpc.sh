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

# sleep
sleep 5

# UpdateUser 
data='{"uid":"=uid","name":"=name","sex":"1"}'
data=$(echo $data | sed s/=uid/$uid/ |sed s/=name/$name/)
grpcurl -plaintext -d $data $ADDR $SERVICE.User/UpdateUser

# sleep
sleep 5

# ReadUser
data='{"uid":"=uid"}'
data=$(echo $data | sed s/=uid/$uid/)
grpcurl -plaintext -d $data $ADDR $SERVICE.User/ReadUser

# sleep
sleep 5

# DeleteUser
grpcurl -plaintext -d $data $ADDR $SERVICE.User/DeleteUser

# sleep
sleep 5

