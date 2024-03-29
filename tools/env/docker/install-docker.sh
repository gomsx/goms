#!/bin/bash
# ubuntu18.04,ubuntu20.04
set -x
set -e
set -u

docker_version=5:20.10.8~3-0

set +x
systemctl stop docker
set -x

sudo apt-get remove docker docker-engine docker-ce docker.io -y
sudo apt-get update -y
sudo apt-get install apt-transport-https ca-certificates curl software-properties-common -y

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
sudo apt-get update -y

sudo apt-cache madison docker-ce
sudo apt-get install docker-ce=$docker_version~ubuntu-$(lsb_release -cs) -y --allow-downgrades

systemctl start docker
systemctl status docker

docker --version
docker version
