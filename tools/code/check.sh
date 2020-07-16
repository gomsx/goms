#!/bin/bash
set -x

# PWD
PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD

# PRO
PRO=$PWD/../..
PRO=$(cd $PRO;pwd)
echo "==> PRO: $PRO"

# chmod +x
find $PRO -name "*.sh" | xargs chmod +x 
# find $PRO -name "*.sh" | xargs -i shellcheck {} 

# doc
$PWD/ck_doc_deta.sh
[ "$1" = "all" ] && $PWD/ck_doc_all.sh

$PWD/ck_doc_deta.sh

# code
$PWD/ck_code_go.sh

