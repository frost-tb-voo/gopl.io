package main

import (
	"testing"

	"golang.org/x/net/html"
)

func TestParsable1(tt *testing.T) {
	ss := Parsable("<html></html>")
	doc, _ := html.Parse(ss)
	for cc := doc.FirstChild; cc != nil; cc = cc.NextSibling {
		tt.Log(cc.Type)
		tt.Log(cc.Data)
	}

	// Output:
	// 3
	// html
}

func TestParsable2(tt *testing.T) {
	ss := Parsable("<html><body></body></html>")
	doc, _ := html.Parse(ss)
	for cc1 := doc.FirstChild; cc1 != nil; cc1 = cc1.NextSibling {
		for cc := cc1.FirstChild; cc != nil; cc = cc.NextSibling {
			tt.Log(cc.Type)
			tt.Log(cc.Data)
		}
	}

	// Output:
	// 3
	// head
	// 3
	// body
}
