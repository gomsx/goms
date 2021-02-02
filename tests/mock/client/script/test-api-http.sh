#!/bin/bash
set -x
set -e

pwdx="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
pkg="$pwdx/pkg.sh"
source "$pkg"

[ $1 ] && iv=$1 || iv=0.1
[ $2 ] && version="/$2" || version=""
[ $3 ] && readtimes=$5 || readtimes=1
[ $4 ] && host=$3 || host=localhost
[ $5 ] && port=$4 || port=8080

FCMD='格式: cmd "间隔时间(单位s)" "版本" "读操作次数" "地址" "http端口"'
ECMD='例子: cmd "0.01" "v1" "1" "127.0.0.1" "8080"'

if [ $# -eq 0 ]; then
	echo "HELP:"
	echo "$FCMD"
	echo "$ECMD"
	exit
fi

addr="$host:$port"
flag=""
# flag="-i -w \"\n\""

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
res="$(curl -X POST -d $data $addr$version/users $flag)"
delay

# get uid
uidx="$(getJsonValueByKey "$res" "uid")"


# get /users
for((i=0; i<readtimes; i++))
do
	curl -X GET $addr$version/users/$uidx $flag
	curl -X GET $addr$version/users?uid=$uidx $flag
	delay
done

# put /users
name=name${uidx:0:5}
data="name=$name&sex=1"
curl -X PUT -d $data $addr$version/users/$uidx $flag
delay

# delete /users
curl -X DELETE $addr$version/users/$uidx $flag
delay
