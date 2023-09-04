#!/bin/sh
sftp 192.168.145.66 <<EOF
cd /usr/local
put mysql-cc.zip
unzip mysql-cc.zip

quit
EOF
