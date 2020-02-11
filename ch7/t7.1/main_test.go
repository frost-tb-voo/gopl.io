package main

import (
	"fmt"
	"testing"
)

func TestWordCounter(tt *testing.T) {
	var c WordCounter
	c.Write([]byte("hello"))
	tt.Log(c)

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	tt.Log(c)

	// Output:
	// 1
	// 2
}

func TestLineCounter(tt *testing.T) {
	var c LineCounter
	c.Write([]byte("hello"))
	tt.Log(c)

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s.\nhow are you?\nthanks.", name)
	tt.Log(c)

	// Output:
	// 1
	// 3
}
