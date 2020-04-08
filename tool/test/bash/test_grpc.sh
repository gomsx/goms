#!/bin/bash
set -x

ADDR=192.168.43.201:50051
	
# ping
# GET /ping
grpcurl -plaintext $ADDR service.goms.User/Ping 

# GET /ping
grpcurl -plaintext -d '{"Message": "xxx"}' $ADDR service.goms.User/Ping 
	