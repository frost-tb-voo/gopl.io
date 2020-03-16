package main

import (
	"fmt"
	"time"
)

func main() {
	chain(10)
}

func chain(length int) {
	src := make(chan struct{})
	in := src
	for ii := 0; ii < length; ii++ {
		out := make(chan struct{})
		go connect(in, out)
		in = out
	}
	dst := in

	start := time.Now()
	src <- struct{}{}
	<-dst
	fmt.Printf("%v, %v\n",
		length, time.Since(start))
}

func connect(in <-chan struct{}, out chan<- struct{}) {
	dat := <-in
	out <- dat
}
