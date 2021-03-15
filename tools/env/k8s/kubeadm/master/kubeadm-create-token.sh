#!/bin/bash
set -x
set -e

# work dir
WD="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "--> work dir: ${WD}"

# output token to file
outfile="${WD}/token"

# work
kubeadm token create | tee ${outfile}

openssl x509 -pubkey -in /etc/kubernetes/pki/ca.crt |
	openssl rsa -pubin -outform der 2>/dev/null |
	openssl dgst -sha256 -hex |
	sed 's/^.* //' |
	tee -a ${outfile}
