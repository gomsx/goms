#!/bin/bash

set -x

VERSION=v1.6.1

curl -sSL -o argocd https://github.com/argoproj/argo-cd/releases/download/$VERSION/argocd-linux-amd64

mv argocd /usr/local/bin/argocd

chmod +x /usr/local/bin/argocd