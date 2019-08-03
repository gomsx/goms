#!/bin/bash
go build ../
sudo chmod +x ./eDocker
sudo docker build -t sfw/edocker -f ./dockerfile ./
sudo docker run -it sfw/edocker
