#!/bin/bash
# set -xe

NS=goms-test

./namespace/create.sh $NS
./configmap/create.sh $NS
./volume/create.sh  $NS
./deployment/apply.sh $NS
./service/apply.sh $NS

