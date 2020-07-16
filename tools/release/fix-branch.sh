#!/bin/bash

set -x

VER="$1"

[ $VER ] || exit

sed -i "s/_master/_$VER/g" README.md
sed -i "s/\/master/\/release-$VER/g" README.md

sed -i "s/_master/_$VER/g" .github/workflows/make_master.yml
sed -i "s/master/release-$VER/g" .github/workflows/make_master.yml
mv .github/workflows/make_master.yml .github/workflows/make_$VER.yml

