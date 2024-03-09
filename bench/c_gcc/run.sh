#!/bin/sh

case `/bin/arch` in
    aarch64)
    CC1=/usr/lib/gcc-cross/x86_64-linux-gnu/11/cc1
    ;;
    riscv64)
    CC1=/usr/lib/gcc-cross/x86_64-linux-gnu/11/cc1
    ;;
    x86_64)
    CC1=/usr/lib/gcc/x86_64-linux-gnu/11/cc1
esac

$CMD_PREFIX $CC1 duktape.i -quiet -O3 >stdout.txt 2>stderr.txt

