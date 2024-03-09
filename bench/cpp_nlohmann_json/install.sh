#! /bin/sh
set -e
if [ -f install_ok ]; then
    exit 0
fi

g++ -O2 -g -std=c++11 test.cpp -o test

touch install_ok
