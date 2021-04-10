#!/bin/bash
set -xe

cd mysql-initdbd && bash delete-vol.sh
bash pv/delete-vol.sh
