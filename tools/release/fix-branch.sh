#!/bin/bash

set -x

REL="v1.7.x"

sed -i "s/_master/_$REL/g" README.md
sed -i "s/\/master/\/release-$REL/g" README.md

sed -i "s/_master/_$REL/g" .github/workflows/make_master.yml
sed -i "s/master/release-$REL/g" .github/workflows/make_master.yml
mv .github/workflows/make_master.yml .github/workflows/make_$REL.yml

