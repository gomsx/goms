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

