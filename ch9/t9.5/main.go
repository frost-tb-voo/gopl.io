package main

import (
	"fmt"
	"time"
)

func main() {
	interact(time.Duration(time.Second * 10))
}

func interact(tt time.Duration) {
	src := make(chan int)
	dst := make(chan int)
	abort := make(chan struct{})
	go connect(src, dst, abort)
	go connect(dst, src, abort)

	src <- 0
	time.Sleep(tt)
	close(abort)
	count := <-dst
	fmt.Printf("%v, %v times\n",
		tt, count)
}

func connect(in <-chan int, out chan<- int, abort <-chan struct{}) {
	for {
		select {
		case <-abort:
			break
		case count := <-in:
			out <- count + 1
		}
	}
}
