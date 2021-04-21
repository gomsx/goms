#!/bin/bash
set -xe

bash mysql-pv/create-vol.sh
bash redis-pv/create-vol.sh
(cd initdbd && bash create-vol.sh)
