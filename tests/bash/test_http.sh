#!/bin/bash
set -e
set -x

[ $1 ] && IV=$1 || IV=0.1
[ $2 ] && VERSION="/$2" || VERSION="" 
[ $3 ] && HOST=$3 || HOST=localhost
[ $4 ] && PORT=$4 || PORT=8080

ADDR="$HOST:$PORT"
FLAG="-i -w \"\n\""

function delay(){
    sleep "$IV"s
    return
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

