#!/bin/bash
set -x

# work dir
WD="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "--> work dir: ${WD}"

# deploy
kubectl apply -f ${WD}/0_14_0/kube-flannel.yml
