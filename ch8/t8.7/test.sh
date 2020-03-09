#!/bin/sh

go get -u golang.org/x/net/html
go run webmirror.go -depth=5 https://golang.org/
