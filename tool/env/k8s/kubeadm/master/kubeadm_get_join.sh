#!/bin/bash
set -x
# set -e

KIP="hostname -i"
TOK="kubeadm token create"
SHA="openssl x509 -pubkey -in /etc/kubernetes/pki/ca.crt | openssl rsa -pubin -outform der 2>/dev/null | openssl dgst -sha256 -hex | sed 's/^.* //'"

MIP=$(eval $KIP)
TOKEN=$(eval $TOK)
SHA256=$(eval $SHA)

set +x

echo "---------------get join arg-----------------------
MIP=$MIP 
TOKEN=$TOKEN 
SHA256=$SHA256"

echo "---------------node join bash---------------------- 
sudo rm -f /etc/kubernetes/admin.conf 
sudo kubeadm reset 
sudo kubeadm join $MIP:6443 --token $TOKEN --discovery-token-ca-cert-hash sha256:$SHA256"

