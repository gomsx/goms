#!/bin/bash
set -x
# set -e

kip="hostname -i"
tok="kubeadm token create"
sha="openssl x509 -pubkey -in /etc/kubernetes/pki/ca.crt | openssl rsa -pubin -outform der 2>/dev/null | openssl dgst -sha256 -hex | sed 's/^.* //'"

mip=$(eval $kip)
token=$(eval $tok)
sha256=$(eval $sha)

set +x

echo "---------------get join arg-----------------------
mip=$mip
token=$token
sha256=$sha256

echo "---------------node join bash---------------------- 
sudo rm -f /etc/kubernetes/admin.conf 
sudo kubeadm reset 
sudo kubeadm join $mip:6443 --token $token --discovery-token-ca-cert-hash sha256:$sha256

