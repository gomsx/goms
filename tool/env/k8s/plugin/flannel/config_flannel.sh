#!/bin/bash
# set -xe

mkdir -p /etc/cni/net.d/

cat <<EOF> /etc/cni/net.d/10-flannel.conf
{"name":"cbr0","type":"flannel","delegate": {"isDefaultGateway": true}}
EOF

rm -rf /usr/share/oci-umount/oci-umount.d
rm -rf /run/flannel/
mkdir /usr/share/oci-umount/oci-umount.d -p
mkdir /run/flannel/

touch /run/flannel/subnet.env
chmod a+rwx /run/flannel/subnet.env

cat <<EOF> /run/flannel/subnet.env
FLANNEL_NETWORK=172.100.0.0/16
FLANNEL_SUBNET=172.100.1.0/24
FLANNEL_MTU=1450
FLANNEL_IPMASQ=true
EOF
