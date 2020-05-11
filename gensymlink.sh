#!/bin/sh

DIR=$(cd $(dirname ${BASH_SOURCE:-$0}); pwd)

mkdir -p ${GOPATH}/src/github.com/frost-tb-voo/
ln -s ${DIR} ${GOPATH}/src/github.com/frost-tb-voo/gopl.io
