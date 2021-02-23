#!/bin/bash
set -x
set -e

# 项目目录
PD="$(cd "../../" && pwd)"
echo "--> pro dir: ${PD}"

# BIN 目录
BIN="${PD}/bin"
[ -d "${BIN}" ] || mkdir "${BIN}"

# PATH
export PATH="${BIN}:$PATH"

# grpcurl
install_grpcurl() {
	wget https://github.com/fullstorydev/grpcurl/releases/download/v1.6.0/grpcurl_1.6.0_linux_x86_64.tar.gz
	tar -C "${BIN}" -xvf grpcurl_1.6.0_linux_x86_64.tar.gz
	rm grpcurl_1.6.0_linux_x86_64.tar.gz
}
which grpcurl || install_grpcurl
