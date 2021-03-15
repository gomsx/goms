#!/bin/bash
# set -xe

# work dir
WD="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "--> work dir: ${WD}"

# work
kubectl apply -f ${WD}/2_0_0/recommended-my.yaml
kubectl apply -f ${WD}/2_0_0/dashboard-adminuser.yaml
