#!/bin/bash
set -x
set -e

## image
./images/k8s-pull-flannel.sh

## plugin
### net
./plugin/flannel/config-flannel.sh
