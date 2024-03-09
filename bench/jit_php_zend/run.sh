#!/bin/sh

$CMD_PREFIX /usr/bin/php -dopcache.enable=1 -dopcache.enable_cli=1 -dopcache.jit_buffer_size=1000M zend_bench.php >stdout.txt 2>stderr.txt

