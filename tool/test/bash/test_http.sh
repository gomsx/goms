#!/bin/bash
set -x

[ $1 ] && IP=$1 || IP=192.168.43.201
[ $2 ] && PORT=$2 || PORT=8080

ADDR="$IP:$PORT"

# ping
# GET /ping
curl -X GET $ADDR/ping -w "\n"

# GET /ping
curl -X GET $ADDR/ping?message=xxx -w "\n"
	
# user
# POST /user/user
# curl -X POST -d "name=xxx&sex=1" $ADDR/user -w "\n"

res=$(curl -X POST -d "name=xxx&sex=1" $ADDR/user);     #echo $res;
res=${res##*\"uid\":};  #echo $res;
res=${res%%\}*};        #echo $res;
uid=$res;
name=name${uid:0:5};    #echo $name

# PUT /user/user
curl -X PUT -d "name=$name&sex=1" $ADDR/user/$uid -w "\n"

# GET /user/user
curl -X GET $ADDR/user/$uid -w "\n"
curl -X GET $ADDR/user?uid=$uid -w "\n"

# DELETE /user/user
curl -X DELETE $ADDR/user/$uid -w "\n"