#!/bin/bash
set -xe

bash pv/create-vol.sh
(cd initdbd && bash create-vol.sh)
