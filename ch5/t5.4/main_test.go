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
	links := visit(nil, doc)
	for _, link := range links {
		tt.Logf("%v\n", link)
	}
}
