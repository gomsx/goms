#!/bin/bash
set -e
set -x

[ $1 ] && IV=$1 || IV=0.1
[ $2 ] && SERVICE="service.goms.$2" ||SERVICE="service.goms" 
[ $3 ] && HOST=$3 || HOST=localhost
[ $4 ] && PORT=$4 || PORT=50051

ADDR="$HOST:$PORT"

function delay(){
    sleep "$IV"s
    return
}

# CreateUser
DATA='{"name":"xxx","sex":"1"}'
CMD="grpcurl -plaintext -d \$DATA \$ADDR \$SERVICE.User/CreateUser"
RES=$(eval $CMD)
delay

RES=$(echo $RES | awk 'NR==1{ print $3 }' | tr -d \"\"\")
UIDX=$RES

# ReadUser
DATA='{"uid":"=uid"}'
DATA=$(echo $DATA | sed s/=uid/$UIDX/)
for I in {1..100};do
    grpcurl -plaintext -d $DATA $ADDR $SERVICE.User/ReadUser
    delay
done

# DeleteUser
DATA='{"uid":"=uid"}'
DATA=$(echo $DATA | sed s/=uid/$UIDX/)
grpcurl -plaintext -d $DATA $ADDR $SERVICE.User/DeleteUser
delay

