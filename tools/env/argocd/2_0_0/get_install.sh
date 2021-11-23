#!/bin/bash
set -x

version=v2.0.0

# wget https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
wget https://raw.githubusercontent.com/argoproj/argo-cd/${version}/manifests/install.yaml

