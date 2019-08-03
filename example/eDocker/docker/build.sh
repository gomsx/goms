#!/bin/bash
#echo $PATH;which go
go build ../
ls -l ./eDocker ../
chmod +x ./eDocker #重要
docker build -t sfw/edocker .
docker run -it sfw/edocker
