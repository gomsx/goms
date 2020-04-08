#!/bin/bash
set -e

PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD

bash $PWD/test_http.sh
bash $PWD/test_grpc.sh