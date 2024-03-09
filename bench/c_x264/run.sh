#!/bin/sh

$CMD_PREFIX /usr/bin/x264 --threads 1 --quiet bench.y4m -o output.mp4 >stdout.txt 2>stderr.txt

