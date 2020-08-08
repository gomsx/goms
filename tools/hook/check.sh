#!/bin/bash
# set -x

echo -e "==> start check ..."

# PWD
PWD=$(cd "$(dirname "$0")";pwd)
# echo "--> PWD: $PWD"

# PRO
PRO=$PWD/../..
PRO=$(cd $PRO;pwd)
# echo "--> PRO: $PRO"

# chmod +x
find $PRO -name "*.sh" | xargs chmod +x 
# find $PRO -name "*.sh" | xargs -i shellcheck {} 

# doc
[ "$1" ] || $PWD/ck-doc-deta.sh
[ "$1" = "all" ] && $PWD/ck-doc-all.sh

# code
$PWD/ck-code-go.sh

echo -e "==> end check"

