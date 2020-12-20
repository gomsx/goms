#!/bin/bash
set -x
set -e

[ $1 ] && iv=$1 || iv=0.1
[ $2 ] && version="/$2" || version=""
[ $3 ] && host=$3 || host=localhost
[ $4 ] && port=$4 || port=8080

addr="$host:$port"
flag="-i -w \"\n\""

function delay() {
	sleep "$iv"s
	return
}

echo "----------ping-----------"
# get /ping
curl -X GET $addr$version/ping $flag
# get /ping
curl -X GET $addr$version/ping?message=xxx $flag

echo "----------user-----------"
# post /users
data="name=xxx&sex=1"
cmd="curl -X POST -d \$data \$addr\$version/users \$flag"
res=$(eval $cmd)
delay

res=${res##*\"uid\":}
res=${res%%\}*}
uidx=$res

# get /users
curl -X GET $addr$version/users/$uidx $flag
curl -X GET $addr$version/users?uid=$uidx $flag
delay

# put /users
name=name${uidx:0:5}
data="name=$name&sex=1"
curl -X PUT -d $data $addr$version/users/$uidx $flag
delay

# delete /users
curl -X DELETE $addr$version/users/$uidx $flag
delay
