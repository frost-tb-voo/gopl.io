#!/bin/sh

go run main.go -format jpeg < ../../ch3/t3.8-float/out-64.png > out-64.jpeg
go run main.go -format gif < ../../ch3/t3.8-float/out-64.png > out-64.gif
