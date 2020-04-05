#!/bin/bash
set -xe
 
sudo apt-get remove docker docker-engine docker-ce docker.io
sudo apt-get update
sudo apt-get install -y apt-transport-https ca-certificates curl software-properties-common  1786  curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
sudo apt-get update

apt-cache madison docker-ce
sudo apt-get install docker-ce=5:18.09.0~3-0~ubuntu-bionic
systemctl status docker