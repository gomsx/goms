#!/bin/bash
set -x
[ $1 ] && SERVICE="service.goms.$1" ||SERVICE="service.goms" 
[ $2 ] && HOST=$2 || HOST=localhost
[ $3 ] && PORT=$3 || PORT=50051

ADDR="$HOST:$PORT"

# CreateUser
data='{"name":"xxx","sex":"1"}'
CMD="grpcurl -plaintext -d \$data \$ADDR \$SERVICE.User/CreateUser"

res=$(eval $CMD)

CMD="echo $res | awk 'NR==1{ print \$3 }' | tr -d \"\"\""
res=$(eval $CMD)
uid=$res

# ReadUser
data='{"uid":"=uid"}'
data=$(echo $data | sed s/=uid/$uid/)
for I in {1..1000};do
    grpcurl -plaintext -d $data $ADDR $SERVICE.User/ReadUser  
done

# DeleteUser
grpcurl -plaintext -d $data $ADDR $SERVICE.User/DeleteUser
