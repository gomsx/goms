#!/bin/bash

set -x
set -e
set -u

./kubeadm/master/kubeadm-get-join.sh

