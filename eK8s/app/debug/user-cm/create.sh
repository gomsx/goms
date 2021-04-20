#!/bin/bash

NS="$1"
[[ -z "$NS" ]] || NS="$(kubens -c)"

kubectl create -n "$NS" cm cm-user --from-file=./configs
