#!/bin/bash
set -xe

bash pv/delete-vol.sh
(cd initdbd && bash delete-vol.sh)
