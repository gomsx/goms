#!/bin/bash
set -x

[ $1 ] && VERSION="/$1" || VERSION="" 
[ $2 ] && HOST=$2 || HOST=localhost
[ $3 ] && PORT=$3 || PORT=8080

ADDR="$HOST:$PORT"

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

# GET /users
curl -X GET $ADDR$VERSION/users/$uid -w "\n"
curl -X GET $ADDR$VERSION/users?uid=$uid -w "\n"

# PUT /users
curl -X PUT -d "name=$name&sex=1" $ADDR$VERSION/users/$uid -w "\n"

# DELETE /users
curl -X DELETE $ADDR$VERSION/users/$uid -w "\n"
