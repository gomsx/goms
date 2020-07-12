#!/bin/bash
set -x

# pwd
PWD=$(cd "$(dirname "$0")";pwd)
echo $PWD

# chmod +x
find $PRO -name "*.sh" | xargs chmod +x 
# find $PRO -name "*.sh" | xargs -i shellcheck {} 

# doc
$PWD/ck_doc_deta.sh
[ "$1" = "all" ] && $PWD/ck_doc_all.sh

$PWD/ck_doc_deta.sh

# code
$PWD/ck_code_go.sh

