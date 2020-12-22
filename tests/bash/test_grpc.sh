#!/bin/bash
set -x
set -e


[ $1 ] && iv=$1 || iv=0.1
[ $2 ] && service="service.goms.$2" ||service="service.goms" 
[ $3 ] && host=$3 || host=localhost
[ $4 ] && port=$4 || port=50051

addr="$host:$port"

function delay(){
    sleep "$iv"s
    return
}

echo "-------------ping---------------"
# Ping
grpcurl -plaintext $addr $service.User/Ping 
# Ping
grpcurl -plaintext -d '{"message":"xxx"}' $addr $service.User/Ping 

echo "-------------user---------------"
# CreateUser
data='{"name":"xxx","sex":"1"}'
cmd="grpcurl -plaintext -d \$data \$addr \$service.User/CreateUser"
res=$(eval $cmd)
delay

res=$(echo $res | awk 'NR==1{ print $3 }' | tr -d \"\"\")
uidx=$res

# ReadUser
data='{"uid":'\"$uidx\"'}'
grpcurl -plaintext -d $data $addr $service.User/ReadUser
delay

# UpdateUser 
name=name${uid:1:6}
data='{"uid":'\"$uidx\"',"name":'\"$name\"',"sex":"1"}'
grpcurl -plaintext -d $data $addr $service.User/UpdateUser
delay

# DeleteUser
data='{"uid":'\"$uidx\"'}'
grpcurl -plaintext -d $data $addr $service.User/DeleteUser
delay

