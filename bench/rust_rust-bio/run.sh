#!/bin/sh

cd rust-bio_bench
$CMD_PREFIX  ./target/release/rust-bio_bench >../stdout.txt 2>../stderr.txt

