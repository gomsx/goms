#!/bin/bash

set -x
set -e
set -u

./kubeadm/master/kubeadm_get_join.sh

