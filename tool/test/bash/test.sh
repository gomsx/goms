#!/bin/bash
set -x

ADDR=192.168.43.201:8080

while true
do
	# 使用 http 方法 GET /ping
	curl -X GET $ADDR/ping -w "\n"

	# 使用 http 方法 GET /ping, 参数 message=xxx
	curl -X GET $ADDR/ping?message=xxx -w "\n"
	
	# 使用 http 方法 POST /user/user, 参数 name=xxx sex=0
	curl -X POST -d "name=xxx&sex=1" $ADDR/user -w "\n"

	res=$(curl -X POST -d "name=xxx&sex=1" 192.168.43.201:8080/user);#echo $res;
	res=${res##*\"uid\":};#echo $res;
	res=${res%%\}*};#echo $res;
	uid=$res;
	name=name${uid:0:5};#echo $name

	# 使用 http 方法 PUT /user/user, 参数 uid=123 name=yyy sex=1
	curl -X PUT -d "name=$name&sex=1" $ADDR/user/$uid -w "\n"

	# 使用 http 方法 GET /user/user, 参数 uid=123
	curl -X GET $ADDR/user/$uid -w "\n"
	# curl -X GET $ADDR/user?uid=$uid -w "\n"

	# 使用 http 方法 DELETE /user/user, 参数 uid=123
	curl -X DELETE $ADDR/user/$uid -w "\n"
done
