package main

import (
	"bufio"
	"bytes"
	"fmt"
)

//!+bytecounter

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewBuffer(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		scanner.Text()
		*c += WordCounter(1)
	}
	return int(*c), nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewBuffer(p))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		scanner.Text()
		*c += LineCounter(1)
	}
	return int(*c), nil
}

//!-bytecounter

func main() {
	//!+main
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")
	//!-main
}
