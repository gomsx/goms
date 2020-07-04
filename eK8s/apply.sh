#!/bin/bash
# set -xe

./namespace/create.sh
./configmap/create.sh
./volume/create.sh
./deployment/apply.sh
./service/apply.sh

