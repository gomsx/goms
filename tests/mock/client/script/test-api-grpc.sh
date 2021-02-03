#!/bin/bash
set -x
set -e

[ $1 ] && iv=$1 || iv=0.1
[ $2 ] && service="service.goms.$2" || service="service.goms"
[ $3 ] && readtimes=$5 || readtimes=1
[ $4 ] && host=$3 || host=localhost
[ $5 ] && port=$4 || port=50051

FCMD='格式: cmd "间隔时间(单位s)" "版本" "读操作次数" "地址" "grpc端口"'
ECMD='例子: "0.01" "v1" "1" "127.0.0.1" "50051"'

if [ $# -eq 0 ]; then
	echo "HELP:"
	echo "$FCMD"
	echo "$ECMD"
	exit
fi

function delay() {
	sleep "$iv"s
	return
}

addr="$host:$port"

version=$2

if [ -z "$version" ]; then
	data_ping='{"message":"xxx"}'
	data_create='{"name":"xxx","sex":"1"}'
	data_read='{"uid":"=uid"}'
	data_update='{"uid":"=uid","name":"=name","sex":"1"}'
	data_delete='{"uid":"=uid"}'
else
	data_ping='{"data":{"message":"xxx"}}'
	data_create='{"data":{"name":"xxx","sex":"1"}}'
	data_read='{"data":{"uid":"=uid"}}'
	data_update='{"data":{"uid":"=uid","name":"=name","sex":"1"}}'
	data_delete='{"data":{"uid":"=uid"}}'
fi

echo "-------------ping---------------"
# Ping
grpcurl -plaintext $addr $service.User/Ping
# Ping
grpcurl -plaintext -d $data_ping $addr $service.User/Ping

echo "-------------user---------------"

# CreateUser
data=$data_create
res="$(grpcurl -plaintext -d $data $addr $service.User/CreateUser)"
delay

# hand res
if [ -z "$version" ]; then
	res="$(echo $res | awk 'NR==1{ print $3 }' | tr -d \"\"\")"
else
	res="$(echo $res | awk 'NR==1{ print $9 }' | tr -d \"\"\")"
fi
uidx=$res

# ReadUser
data=$(echo $data_read | sed s/=uid/$uidx/)

for((i=0; i<readtimes; i++))
do
	grpcurl -plaintext -d $data $addr $service.User/ReadUser
	delay
done

# UpdateUser
name="name${uid:1:6}"
data=$(echo $data_update | sed s/=uid/$uidx/ | sed s/=name/$name/)
grpcurl -plaintext -d $data $addr $service.User/UpdateUser
delay

# DeleteUser
data=$(echo $data_delete | sed s/=uid/$uidx/)
grpcurl -plaintext -d $data $addr $service.User/DeleteUser
delay
