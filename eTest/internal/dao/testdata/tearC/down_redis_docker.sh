#!/bin/bash
set -x
# set -e

#rm docker
docker stop redistest 
docker rm redistest 

#ps docker
docker ps | grep redistest

