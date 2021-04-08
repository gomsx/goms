#!/bin/bash
set -x
set -e

version=v1.7.8

wget -O argocd https://github.com/argoproj/argo-cd/releases/download/${version}/argocd-linux-amd64

mv argocd /usr/local/bin/argocd
chmod +x /usr/local/bin/argocd
