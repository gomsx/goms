#!/bin/bash
set -x

[ $1 ] && IP=$1 || IP=192.168.43.201
[ $2 ] && PORT=$2 || PORT=50051

ADDR="$IP:$PORT"
	
# ping
# Ping
grpcurl -plaintext $ADDR service.goms.User/Ping 

# Ping
grpcurl -plaintext -d '{"Message": "xxx"}' $ADDR service.goms.User/Ping 

# user
# CreateUser
res=$(grpcurl -plaintext -d '{"Name": "xxx","Sex":"0"}' $ADDR service.goms.User/CreateUser)
res=$(echo $res | awk 'NR==1{ print $3 }' | tr -d "\"")
uid=$res;
name=name${uid:1:6};echo $name

# UpdateUser 
data='{"Uid":"uid","Name":"name","Sex":"1"}'
data=$(echo $data | sed s/uid/$uid/ |sed s/name/$name/)
grpcurl -plaintext -d $data $ADDR service.goms.User/UpdateUser

# ReadUser
data='{"Uid":"uid"}'
data=$(echo $data | sed s/uid/$uid/)
grpcurl -plaintext -d $data $ADDR service.goms.User/ReadUser

# DeleteUser
grpcurl -plaintext -d $data $ADDR service.goms.User/DeleteUser