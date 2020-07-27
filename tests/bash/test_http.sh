#!/bin/bash
set -x

[ $1 ] && US=$1 || US=100
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

# ping
# GET /ping
curl -X GET $ADDR$VERSION/ping -w "\n"

# GET /ping
curl -X GET $ADDR$VERSION/ping?message=xxx -w "\n"

# user
# POST /users
res=$(curl -X POST -d "name=xxx&sex=1" $ADDR$VERSION/users); 
res=${res##*\"uid\":};  
res=${res%%\}*};        
uid=$res;
name=name${uid:0:5};   

delay

# GET /users
curl -X GET $ADDR$VERSION/users/$uid -w "\n"
curl -X GET $ADDR$VERSION/users?uid=$uid -w "\n"

delay

# PUT /users
curl -X PUT -d "name=$name&sex=1" $ADDR$VERSION/users/$uid -w "\n"

delay

# DELETE /users
curl -X DELETE $ADDR$VERSION/users/$uid -w "\n"

delay