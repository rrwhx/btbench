#!/bin/sh

cd book
$CMD_PREFIX  ../mdBook-0.4.30/target/release/mdbook build >../stdout.txt 2>../stderr.txt

