package main

import (
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestVisit(tt *testing.T) {
	ff, err := os.Open("index.html")
	if err != nil {
		tt.Fatalf("%v\t", err)
		tt.Fail()
	}
	doc, err := html.Parse(ff)
	if err != nil {
		tt.Fatalf("%v\t", err)
		tt.Fail()
	}
	out := visit(map[string]int{}, doc)
	for key, value := range out {
		tt.Logf("%v:\t%v\n", key, value)
	}
}
