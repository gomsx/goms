#!/bin/bash
set +x

[ $1 ] && US=$1 || US=10
[ $2 ] && SERVICE="service.goms.$2" ||SERVICE="service.goms" 
[ $3 ] && HOST=$3 || HOST=localhost
[ $4 ] && PORT=$4 || PORT=50051

ADDR="$HOST:$PORT"

function delay(){
    for ((i=0;i<"$US";i="$i"+1))
    do
        # sleep 0.01
        a=1
    done
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

