#!/bin/bash

docker pull coredns/coredns:1.8.4
docker tag docker.io/coredns/coredns:1.8.4 registry.aliyuncs.com/google_containers/coredns:v1.8.4
docker rmi docker.io/coredns/coredns:1.8.4
