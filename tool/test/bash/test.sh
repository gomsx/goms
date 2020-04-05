#!/bin/bash

while true
do
	curl -X GET 192.168.43.201:31003/ping -w "\n"
	curl -X GET 192.168.43.201:8080/ping?message=xxx -w "\n"
	curl -X POST -d "name=xxx&sex=1"  192.168.43.201:31003/user -w "\n"
	curl -X POST -d "name=xxx&sex=1"  192.168.43.201:8080/user -w "\n"
	curl -X GET 192.168.43.204:31003/ping -w "\n"
	curl -X GET 192.168.43.204:8080/ping?message=xxx -w "\n"
	curl -X POST -d "name=xxx&sex=1"  192.168.43.204:31003/user -w "\n"
	curl -X POST -d "name=xxx&sex=1"  192.168.43.204:8080/user -w "\n"
done
