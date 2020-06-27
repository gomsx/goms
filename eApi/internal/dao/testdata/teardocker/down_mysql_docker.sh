#!/bin/bash
set -x
# set -e

#rm docker
docker stop mysqltest 
docker rm mysqltest 

#ps docker
docker ps | grep mysqltest

