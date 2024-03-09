#!/bin/sh
set -e
if [ -f install_ok ]; then
    exit 0
fi

cd bench
go get -u github.com/go-ego/gse
go build -buildvcs=false
cd -
touch install_ok

