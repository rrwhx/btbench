FROM ubuntu:22.04 AS base
ARG DEBIAN_FRONTEND=noninteractive
RUN sed -i 's@//.*ubuntu.com@//mirrors.ustc.edu.cn@g' /etc/apt/sources.list && apt-get update && apt-get install -y apt-utils
RUN yes | unminimize 
RUN apt-get -q install -y vim cmake curl file git make time wget tar gzip xz-utils x264 gcc g++ nodejs golang php8.1 python3-minimal pypy xalan
RUN curl https://sh.rustup.rs -sSf | bash -s -- -y
ENV PATH="/root/.cargo/bin:${PATH}"
RUN apt-get clean
ADD . /root/btbench
WORKDIR /root
