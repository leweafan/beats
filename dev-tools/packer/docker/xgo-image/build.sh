#!/bin/sh

docker pull tudorg/xgo-base:v20180222 && \
    docker build --rm=true -t tudorg/xgo-1.7.6 go-1.7.6/ &&
    docker build --rm=true -t tudorg/beats-builder beats-builder
