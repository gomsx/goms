#!/bin/bash
set -x
set -e

# 项目 root
PD="$(cd "../../" && pwd)"
BIN="${PD}/bin"
[ -d "${BIN}" ] || mkdir "${BIN}"
export PATH="${BIN}:$PATH"

install_grpcurl() {
	wget https://github.com/fullstorydev/grpcurl/releases/download/v1.6.0/grpcurl_1.6.0_linux_x86_64.tar.gz
	tar -C "${BIN}" -xvf grpcurl_1.6.0_linux_x86_64.tar.gz
	rm grpcurl_1.6.0_linux_x86_64.tar.gz
}

which grpcurl || install_grpcurl
