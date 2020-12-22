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

# post /users
cmd="curl -X POST -d \$data \$addr\$version/users \$flag"
res=$(eval $cmd)

res=${res##*\"uid\":}
res=${res%%\}*}
uidx=$res

# get /users
for i in {1..100}; do
	curl -X GET $addr$version/users?uid=$uidx $flag
	delay
done

# delete /users
curl -X DELETE $addr$version/users/$uidx $flag