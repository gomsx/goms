#!/bin/bash
# set -xe


sudo kubectl apply -f 2_0_0/recommended-my.yaml

sudo kubectl get rs,pod,deploy,svc,ep -n kubernetes-dashboard

# sudo kubectl delete -f 2_0_0/recommended-my.yaml