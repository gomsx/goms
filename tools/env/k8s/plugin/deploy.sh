#!/bin/bash
set -x
set -e

## addons
### net
bash flannel/deploy-flannel.sh

### dashboard
bash dashboard/deploy-dashboard.sh
