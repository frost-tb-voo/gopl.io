package main

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"fmt"
	"net/http/httptest"

	"golang.org/x/net/html"
)

func TestDatabase(tt *testing.T) {
	req := httptest.NewRequest("POST", "http://example.com/list", nil)
	w := httptest.NewRecorder()

	database := map[string]dollars{"shoes": 50, "socks": 5}
	db := Database{database: database}
	http.HandleFunc("/update", db.update)
	db.list(w, req)
	// db.price(w, req)
	// db.update(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		tt.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode && len(strings.TrimSpace(n.Data)) > 0 {
			fmt.Printf("%v\n", n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	// Output:
	// prices
	// Item
	// Price
	// shoes
	// $50.00
	// socks
	// $5.00
}
