#!/bin/bash
# set -x
# set -e

# https://github.com/containernetworking/plugins/tree/master/plugins/meta/flannel

rm -rf /etc/cni/net.d/
mkdir -p /etc/cni/net.d/

# 10-flannel.conf 名中 10 和 flannel 版本无关，0.8 版中也是就是这名
cat <<EOF> /etc/cni/net.d/10-flannel.conf
{"name":"cbr0","type":"flannel","delegate": {"isDefaultGateway": true}}
EOF

rm -rf /usr/share/oci-umount/oci-umount.d
rm -rf /run/flannel/
mkdir /usr/share/oci-umount/oci-umount.d -p
mkdir /run/flannel/

touch /run/flannel/subnet.env
chmod a+rwx /run/flannel/subnet.env

# ok
cat <<EOF> /run/flannel/subnet.env
FLANNEL_NETWORK=10.1.0.0/16
FLANNEL_SUBNET=10.1.17.1/24
FLANNEL_MTU=1472
FLANNEL_IPMASQ=true
EOF

# cat <<EOF> /run/flannel/subnet.env
# FLANNEL_NETWORK=172.100.0.0/16
# FLANNEL_SUBNET=172.100.1.0/24
# FLANNEL_MTU=1450
# FLANNEL_IPMASQ=true
# EOF

