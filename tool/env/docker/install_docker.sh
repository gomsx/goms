#!/bin/bash
set -x
set -e
set -u
 
sudo apt-get remove docker docker-engine docker-ce docker.io -y
sudo apt-get update -y
sudo apt-get install apt-transport-https ca-certificates curl software-properties-common  -y

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
sudo apt-get update -y

apt-cache madison docker-ce
sudo apt-get install docker-ce=5:18.09.0~3-0~ubuntu-bionic -y
systemctl status docker