#!/bin/sh

go get -u gopl.io/ch5/links
go run findlinks.go -depth=3 https://golang.org/
