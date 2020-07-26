#!/bin/bash
set -x

[ $1 ] && US=$1 || US=10
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
for I in {1..100};do
    grpcurl -plaintext -d $data $ADDR $SERVICE.User/ReadUser  
    delay
done

# DeleteUser
grpcurl -plaintext -d $data $ADDR $SERVICE.User/DeleteUser
