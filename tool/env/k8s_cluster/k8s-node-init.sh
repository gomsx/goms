#!/bin/bash

set -x
set -e
set -u

## images
./images/k8s_pull_node.sh
./images/k8s_pull_flannel.sh

## addons
### net
./plugin/flannel/config_flannel.sh

