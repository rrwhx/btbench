#/bin/bash
set -x
export DEBIAN_FRONTEND=noninteractive
apt-get -q -y update
apt-get -q install -y vim cmake curl file git make time wget tar gzip xz-utils x264 gcc g++ nodejs golang php8.1 python3-minimal pypy xalan
if [ `arch` != "x86_64" ]
then
    apt-get -q install -y gcc-x86-64-linux-gnu
fi
curl https://sh.rustup.rs -sSf | bash -s -- -y
echo 'source $HOME/.cargo/env' >> $HOME/.bashrc
export PATH="${HOME}/.cargo/bin:${PATH}"
apt-get clean
