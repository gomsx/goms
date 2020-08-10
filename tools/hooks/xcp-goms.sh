#!/bin/bash
# set -x

# dir
dir="$1"

[ $dir ] || exit

# pwd
pwd=$(pwd)
echo "==> pwd: $pwd"

# dst src
dst=$pwd/$dir
src=${dst/fuwensun/vuca}

# copy
rm -rf "$dst"
cp -r "$src" "$dst"


# PWD 
PWD=$(cd "$(dirname "$0")";pwd)
echo "==> PWD: $PWD"

$PWD/xcheck.sh