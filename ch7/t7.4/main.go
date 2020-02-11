package main

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Parsable string

var reader io.Reader

func (pp Parsable) Read(bb []byte) (int, error) {
	if reader == nil {
		reader = strings.NewReader(string(pp))
	}
	return reader.Read(bb)
}

func main() {
	ss := Parsable("<html></html>")
	html.Parse(ss)
}
