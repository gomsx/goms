#!/bin/bash
set -xe

kubectl delete namespace goms-dev
kubectl delete namespace goms-test
kubectl delete namespace goms-prod

