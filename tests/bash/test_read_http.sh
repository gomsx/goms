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

# POST /users
DATA="name=xxx&sex=1"
CMD="curl -X POST -d \$DATA \$ADDR\$VERSION/users \$FLAG"
RES=$(eval $CMD)
delay

RES=${RES##*\"uid\":}; 
RES=${RES%%\}*}      
UIDX=$RES

# GET /users
for I in {1..100};do
    curl -X GET $ADDR$VERSION/users?uid=$UIDX $FLAG
    delay
done

# DELETE /users
curl -X DELETE $ADDR$VERSION/users/$UIDX $FLAG
delay

