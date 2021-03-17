#!/bin/bash
set -x
set -e

# parse arg
[ $1 ] && iv=$1 || iv=0.1
[ $2 ] && service="service.goms.$2" || service="service.goms"
[ $3 ] && readtimes=$3 || readtimes=1
[ $4 ] && host=$4 || host=localhost
[ $5 ] && port=$5 || port=50051
[ $6 ] && set +x

if [ $# -eq 0 ]; then
	FCMD='格式: cmd "间隔时间(单位s)" "版本" "读操作次数" "地址" "grpc端口"'
	ECMD='例子: cmd "0.01" "v1" "1" "127.0.0.1" "50051"'
	echo "HELP:"
	echo "${FCMD}"
	echo "${ECMD}"
	exit 255
fi

. pkg.sh

# work
addr="${host}:${port}"
version=$2

function delay() {
	sleep ${iv}s
	return
}

if [ -z "${version}" ]; then
	data_ping='{"message":"msg-xxx"}'
	data_create='{"name":"=name","sex":"1"}'
	data_read='{"uid":=uid}'
	data_update='{"uid":=uid,"name":"=name","sex":"1"}'
	data_delete='{"uid":=uid}'
else
	data_ping='{"data":{"message":"msg-xxx"}}'
	data_create='{"data":{"name":"=name","sex":"1"}}'
	data_read='{"data":{"uid":=uid}}'
	data_update='{"data":{"uid":=uid,"name":"=name","sex":"1"}}'
	data_delete='{"data":{"uid":=uid}}'
fi

echo "--> ping"
# Ping
grpcurl -plaintext ${addr} ${service}.User/Ping
# Ping
grpcurl -plaintext -d ${data_ping} ${addr} ${service}.User/Ping

echo "--> user"
# CreateUser
data=${data_create}
res="$(grpcurl -plaintext -d ${data} ${addr} ${service}.User/CreateUser)"
# TODO
delay

# get uid
if [ -n "${version}" ]; then
	res="$(getJsonValueByKey "${res}" "data")"
fi
uidx="$(getJsonValueByKey "${res}" "uid")"
[ "failed" == ${uidx} ] && exit 1


# ReadUser
data=$(echo ${data_read} | sed s/=uid/${uidx}/)

for((i=0; i<readtimes; i++))
do
	grpcurl -plaintext -d ${data} ${addr} ${service}.User/ReadUser
	# TODO
	delay
done

# UpdateUser
name="name${uid:1:6}"
data=$(echo ${data_update} | sed s/=uid/${uidx}/ | sed s/=name/${name}/)
grpcurl -plaintext -d ${data} ${addr} ${service}.User/UpdateUser
# TODO
delay

# DeleteUser
data=$(echo ${data_delete} | sed s/=uid/${uidx}/)
grpcurl -plaintext -d ${data} ${addr} ${service}.User/DeleteUser
# TODO
delay
