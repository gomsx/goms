#!/bin/bash

set -x

VER="$1"

[ $VER ] || exit

sed -i "s/_master/_$VER/g" README.md
sed -i "s/\/master/\/release-$VER/g" README.md

sed -i "s/_master/_$VER/g" .github/workflows/make_master.yaml
sed -i "s/master/release-$VER/g" .github/workflows/make_master.yaml
mv .github/workflows/make_master.yaml .github/workflows/make_$VER.yaml

