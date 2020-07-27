#!/bin/bash
set +x

[ $1 ] && US=$1 || US=10
[ $2 ] && VERSION="/$2" || VERSION="" 
[ $3 ] && HOST=$3 || HOST=localhost
[ $4 ] && PORT=$4 || PORT=8080

ADDR="$HOST:$PORT"

function delay(){
    for ((i=0;i<"$US";i="$i"+1))
    do
        # sleep 0.01
        a=1
    done
    echo "==> delay $US us"
}

# POST /users
res=$(curl -X POST -d "name=xxx&sex=1" $ADDR$VERSION/users); 
res=${res##*\"uid\":};  
res=${res%%\}*};        
uid=$res;

# GET /users
for I in {1..100};do
    curl -X GET $ADDR$VERSION/users/$uid -w "\n"
    delay
done

# DELETE /users
curl -X DELETE $ADDR$VERSION/users/$uid -w "\n"

