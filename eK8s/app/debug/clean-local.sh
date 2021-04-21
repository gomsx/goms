#!/bin/bash
set -xe

bash mysql-pv/delete-vol.sh
bash redis-pv/delete-vol.sh
(cd initdbd && bash delete-vol.sh)
