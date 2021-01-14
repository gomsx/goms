#!/bin/bash
set -x
set -e

pwdx=$(
	cd "$(dirname "$0")"
	pwd
)

while true; do
	bash $pwdx/test_read_http.sh "$1" "$2" "$3" "$4"
	bash $pwdx/test_read_grpc.sh "$1" "$2" "$3" "$5"
done
