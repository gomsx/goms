#!/bin/bash
set -x

# 当前目录路径
pwdx="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# 当前项目路径
prox="$(cd "$pwdx/../.." && pwd)"
binx="$prox/bin"
[ -d "$binx" ] || mkdir "$binx"
export PATH="$binx:$PATH"

# install grpc
wget https://github.com/fullstorydev/grpcurl/releases/download/v1.6.0/grpcurl_1.6.0_linux_x86_64.tar.gz
tar -C "$binx" -xvf grpcurl_1.6.0_linux_x86_64.tar.gz
rm grpcurl_1.6.0_linux_x86_64.tar.gz
