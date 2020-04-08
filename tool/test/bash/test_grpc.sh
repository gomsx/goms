#!/bin/bash
set -x

ADDR=192.168.43.201:50051

while true
do
	# ping
	# GET /ping
	 grpcurl -plaintext $ADDR service.goms.User/Ping 

	# GET /ping
	 grpcurl -plaintext -d '{"Message": "xxx"}' $ADDR service.goms.User/Ping 
	
	# user
	# POST /user/user
	#  grpcurl -X POST -d "name=xxx & sex=1" $ADDR/user 

	# res=$( grpcurl -X POST -d "name=xxx & sex=1" $ADDR/user);#echo $res;
	# res=${res##*\"uid\":};#echo $res;
	# res=${res%%\}*};#echo $res;
	# uid=$res;
	# name=name${uid:0:5};#echo $name

	# # PUT /user/user
	#  grpcurl -X PUT -d "name=$name & sex=1" $ADDR/user/$uid 

	# # GET /user/user
	#  grpcurl -X GET $ADDR/user/$uid 
	#  grpcurl -X GET $ADDR/user?uid=$uid 

	# # DELETE /user/user
	#  grpcurl -X DELETE $ADDR/user/$uid 
done
