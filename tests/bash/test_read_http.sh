#!/bin/bash
set -x

[ $1 ] && VERSION="/$1" || VERSION="" 
[ $2 ] && HOST=$2 || HOST=localhost
[ $3 ] && PORT=$3 || PORT=8080

ADDR="$HOST:$PORT"

# POST /users
res=$(curl -X POST -d "name=xxx&sex=1" $ADDR$VERSION/users); 
res=${res##*\"uid\":};  
res=${res%%\}*};        
uid=$res;

# GET /users
for I in {1..1000};do
    curl -X GET $ADDR$VERSION/users/$uid -w "\n"
    curl -X GET $ADDR$VERSION/users?uid=$uid -w "\n"
done

# DELETE /users
curl -X DELETE $ADDR$VERSION/users/$uid -w "\n"
