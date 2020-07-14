#!/bin/bash

set -x

REL="v1.7.x"

sed -i "s/master/$REL/g" README.md
sed -i "s/master/$REL/g" README.md .github/workflows/make_master.yml
