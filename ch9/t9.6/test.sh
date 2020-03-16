#!/bin/sh

GOMAXPROCS=1 go test -v -race *.go > 1.log
GOMAXPROCS=10 go test -v -race *.go > 10.log
GOMAXPROCS=100 go test -v -race *.go > 100.log
GOMAXPROCS=1000 go test -v -race *.go > 1000.log
go test -v -race *.go > log
