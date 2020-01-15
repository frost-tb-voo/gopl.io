package main

import (
	"net/http"
	"testing"

	"golang.org/x/net/html"
)

func TestElementByID(tt *testing.T) {
	url := "https://www.google.com/"
	id1 := "mngb"
	id2 := "gbar"
	id3 := "hoge"
	resp, err := http.Get(url)
	if err != nil {
		tt.Fatalf("%v\n", err)
		tt.Fail()
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		tt.Fatalf("%v\n", err)
		tt.Fail()
	}

	target1 := ElementByID(doc, id1)
	tt.Logf("%v\n", summary(target1))
	target2 := ElementByID(doc, id2)
	tt.Logf("%v\n", summary(target2))
	target3 := ElementByID(doc, id3)
	tt.Logf("%v\n", summary(target3))
}
