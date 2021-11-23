#!/bin/bash
set -x # set -e
set -u

# ctop
sudo apt install ctop

# jp
go get -u github.com/sgreben/jp/cmd/jp