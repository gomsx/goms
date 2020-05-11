#!/bin/bash

set -x
set -e
set -u

PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD

DK=$PWD/../docker
K8S=$PWD/../k8s_cluster

$DK/install_docker.sh
$DK/config_docker.sh



