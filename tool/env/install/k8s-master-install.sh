#!/bin/bash

set -x
set -e
set -u

PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD

$PWD/../docker/install_docker.sh
$PWD/../docker/config_docker.sh


