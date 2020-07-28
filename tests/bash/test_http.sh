#!/bin/bash
# set -x
set +x

[ $1 ] && US=$1 || US=100
[ $2 ] && VERSION="/$2" || VERSION="" 
[ $3 ] && HOST=$3 || HOST=localhost
[ $4 ] && PORT=$4 || PORT=8080

ADDR="$HOST:$PORT"
FLAG="-i -w \"\n\""

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

echo "----------ping-----------"
# GET /ping
curl -X GET $ADDR$VERSION/ping $FLAG
# GET /ping
curl -X GET $ADDR$VERSION/ping?message=xxx $FLAG

echo "----------user-----------"
# POST /users
DATA="name=xxx&sex=1"
CMD="curl -X POST -d \$DATA \$ADDR\$VERSION/users \$FLAG"
RES=$(eval $CMD)
delay

RES=${RES##*\"uid\":}; 
RES=${RES%%\}*}      
UIDX=$RES

# GET /users
curl -X GET $ADDR$VERSION/users/$UIDX $FLAG
curl -X GET $ADDR$VERSION/users?uid=$UIDX $FLAG
delay

# PUT /users
NAME=name${UIDX:0:5} 
DATA="name=$NAME&sex=1"
curl -X PUT -d $DATA $ADDR$VERSION/users/$UIDX $FLAG
delay

# DELETE /users
curl -X DELETE $ADDR$VERSION/users/$UIDX $FLAG
delay
