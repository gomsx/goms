#!/bin/bash
set -x

# flux
git clone https://github.com/fluxcd/flux

sed -i "s/--git-url=git@github.com:fluxcd/flux-get-started/\$s1/g" flux/deploy/flux-deployment.yaml

fluxctl identity --k8s-fwd-ns=flux
fluxctl sync --k8s-fwd-ns=flux

