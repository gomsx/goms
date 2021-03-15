#!/bin/bash
# set -xe

# work dir
WD="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "--> work dir: ${WD}"

# deploy
kubectl apply -f ${WD}/1_2_0/kube-flannel.yml
