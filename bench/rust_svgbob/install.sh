#!/bin/sh
set -e
if [ -f install_ok ]; then
    exit 0
fi
tar -xf svgbob_0.5.5.tar.gz
cd svgbob-0.5.5/
sed -i "s/input_path.clone()/input_path/" svgbob_cli/src/main.rs
cargo build --release
cd -
touch install_ok
