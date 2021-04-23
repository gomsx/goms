#!/bin/bash
set -x
set -e

bash kubeadm/master/kubeadm-gen-join-cmd.sh
