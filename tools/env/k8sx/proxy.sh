#!/bin/bash

# https://www.liuyixiang.com/post/108262.html
# option2：kube-proxy 端口转发
kubectl proxy --address='100.64.198.131' --port=9090 --accept-hosts='^*$'

#URL：http://100.64.198.131:9090/api/v1/namespaces/istio-system/services/http:prometheus:9090/proxy

