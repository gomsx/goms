#!/bin/bash
set -x
set -e

kubeadm token create --print-join-command
