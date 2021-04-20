#!/bin/bash
set -xe

mkdir -p /var/lib/mysqlx/initdbd-debug
chmod 777 /var/lib/mysqlx/initdbd-debug

cp -r setup /var/lib/mysqlx/initdbd-debug
