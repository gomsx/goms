#!/bin/bash
set -x

# pwd
PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD

# PRO
PRO=$PWD/../..

# chmod *.sh
find $PRO -name "*.sh" | xargs chmod +x 

# go
go fmt $PRO/...

# -l 防止内敛
# go build -gcflags '-m -l'