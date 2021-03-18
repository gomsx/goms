#!/bin/bash
set -x
set -e

./kubeadm/master/kubeadm-gen-join-cmd.sh
