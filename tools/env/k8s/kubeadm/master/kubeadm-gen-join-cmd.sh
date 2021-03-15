#!/bin/bash
# set -x
# set -e

# work dir
WD="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "--> work dir: ${WD}"

# get args
master_ip="$(sed -n "s/\(master_ip:\)\(.*\)/\2/p" ${WD}/config.yaml | tr -d ' ')"
token="$(cat ${WD}/token | head -1 | tr -d ' ')"
cert_hash="$(cat ${WD}/token | head -2 | tail -1 | tr -d ' ')"

# print cmd
echo "--> generate join cmd for mode:"
echo "sudo kubeadm reset"
echo "sudo kubeadm join ${master_ip}:6443 --token ${token} --discovery-token-ca-cert-hash sha256:${cert_hash}"
