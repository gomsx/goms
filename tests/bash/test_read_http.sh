#!/bin/bash
set -x

[ $1 ] && US=$1 || US=10
[ $2 ] && VERSION="/$2" || VERSION="" 
[ $3 ] && HOST=$3 || HOST=localhost
[ $4 ] && PORT=$4 || PORT=8080

ADDR="$HOST:$PORT"

# usleep : 默认以微秒。  
# 1s = 1000ms = 1000000us
function delay(){
    # sleep 1
    usleep $US
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

