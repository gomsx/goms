#!/bin/bash
set -xe

cd mysql-initdbd && bash create-vol.sh
bash pv/create-vol.sh
