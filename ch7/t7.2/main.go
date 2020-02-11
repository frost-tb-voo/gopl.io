// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 173.

// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
package main

import (
	"fmt"
	"io"
	"os"
)

type ByteCounterWriter struct {
	counter *int64
	w       *io.Writer
}

func (c ByteCounterWriter) Write(p []byte) (int, error) {
	n, err := (*c.w).Write(p)
	// fmt.Fprintf(os.Stderr, "%v+=%v\n", *c.counter, n)
	*(c.counter) += int64(n)
	// fmt.Fprintf(os.Stderr, "counter:%v\n", *c.counter)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var counter int64
	bcw := ByteCounterWriter{counter: &counter, w: &w}
	return bcw, &counter
}

func main() {
	//!+main
	c, count := CountingWriter(os.Stdout)
	c.Write([]byte("hello"))
	fmt.Println()
	fmt.Println(*count) // "5", = len("hello")

	c, count = CountingWriter(os.Stdout)
	var name = "Dolly"
	fmt.Fprintf(c, "hello, %s", name)
	fmt.Println()
	fmt.Println(*count) // "12", = len("hello, Dolly")
	//!-main
}
