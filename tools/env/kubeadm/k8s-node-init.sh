#!/bin/bash

set -x
set -e
set -u

## images
./images/k8s-pull-node.sh
./images/k8s-pull-flannel.sh

## addons
### net
./plugin/flannel/config-flannel.sh

