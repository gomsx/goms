#!/bin/bash
set -e
set -x

[ $1 ] && iv=$1 || iv=0.1
[ $2 ] && service="service.goms.$2" ||service="service.goms" 
[ $3 ] && host=$3 || host=localhost
[ $4 ] && port=$4 || port=50051

addr="$host:$port"

function delay(){
    sleep "$iv"s
    return
}

# CreateUser
data='{"name":"xxx","sex":"1"}'
cmd="grpcurl -plaintext -d \$data \$addr \$service.User/CreateUser"
res=$(eval $cmd)

res=$(echo $res | awk 'NR==1{ print $3 }' | tr -d \"\"\")
uidx=$res

# ReadUser
data='{"uid":'\"$uidx\"'}'
for I in {1..100};do
    grpcurl -plaintext -d $data $addr $service.User/ReadUser
    delay
done

# DeleteUser
grpcurl -plaintext -d $data $addr $service.User/DeleteUser
