#!/bin/sh
set -e
if [ -f install_ok ]; then
    exit 0
fi

cd rust-bio_bench
cargo build --release

cd -
touch install_ok
