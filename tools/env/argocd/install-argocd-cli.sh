#!/bin/bash
set -x
set -e

version=v1.6.1

curl -ssl -o argocd https://github.com/argoproj/argo-cd/releases/download/$version/argocd-linux-amd64

mv argocd /usr/local/bin/argocd
chmod +x /usr/local/bin/argocd

