#!/bin/bash
# set -x
set +x

[ $1 ] && US=$1 || US=100
[ $2 ] && SERVICE="service.goms.$2" ||SERVICE="service.goms" 
[ $3 ] && HOST=$3 || HOST=localhost
[ $4 ] && PORT=$4 || PORT=50051

ADDR="$HOST:$PORT"

function delay(){
# set +x
    for ((i=0;i<"$US";i="$i"+1))
    do
        # sleep 0.01
        a=1
    done
    echo "==> delay $US us"
# set -x
}

echo "-------------ping---------------"
# Ping
grpcurl -plaintext $ADDR $SERVICE.User/Ping 
# Ping
grpcurl -plaintext -d '{"message":"xxx"}' $ADDR $SERVICE.User/Ping 

echo "-------------user---------------"
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
grpcurl -plaintext -d $DATA $ADDR $SERVICE.User/ReadUser
delay

# UpdateUser 
NAME=name${uid:1:6}
DATA='{"uid":"=uid","name":"=name","sex":"1"}'
DATA=$(echo $DATA | sed s/=uid/$UIDX/ |sed s/=name/$NAME/)
grpcurl -plaintext -d $DATA $ADDR $SERVICE.User/UpdateUser
delay

# DeleteUser
DATA='{"uid":"=uid"}'
DATA=$(echo $DATA | sed s/=uid/$UIDX/)
grpcurl -plaintext -d $DATA $ADDR $SERVICE.User/DeleteUser
delay

