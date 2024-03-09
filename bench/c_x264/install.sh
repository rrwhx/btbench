#!/bin/sh
set -e
if [ -f install_ok ]; then
    exit 0
fi

xz -T 0 -k -d -f bench.y4m.xz

touch install_ok

