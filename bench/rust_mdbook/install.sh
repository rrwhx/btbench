#!/bin/sh

set -e
if [ -f install_ok ]; then
    exit 0
fi

rm -rf book/book/
tar -xf mdbook_v0.4.30.tar.gz
sed -i '20s/content/\&content[0..1]/' mdBook-0.4.30/src/utils/fs.rs
cd mdBook-0.4.30
cargo build --release

cd -
touch install_ok

