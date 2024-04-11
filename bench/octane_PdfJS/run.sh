#!/bin/sh

CUR_D=$(basename "$PWD")
$CMD_PREFIX /usr/bin/node --single-threaded ../octane/octane.js ${CUR_D#*_} >stdout.txt 2>stderr.txt

