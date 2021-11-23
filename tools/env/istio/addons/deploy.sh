#!/bin/bash
set -x

version="1.11.2"

kubectl apply -f /usr/local/istio-${version}/samples/addons/kiali.yaml
kubectl apply -f /usr/local/istio-${version}/samples/addons/grafana.yaml
kubectl apply -f /usr/local/istio-${version}/samples/addons/prometheus.yaml
