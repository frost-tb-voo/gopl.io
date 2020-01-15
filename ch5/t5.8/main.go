// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stdout, "usage: command $url $element_id\n")
		return
	}
	resp, err := http.Get(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	for _, id := range os.Args[2:] {
		target := ElementByID(doc, id)
		fmt.Fprintf(os.Stdout, "%v\n", summary(target))
	}
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) bool {
	if pre != nil {
		if !pre(n) {
			return false
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if !forEachNode(c, pre, post) {
			return false
		}
	}

	if post != nil {
		if !post(n) {
			return false
		}
	}
	return true
}

//!-forEachNode

// ElementByID is ..
func ElementByID(doc *html.Node, id string) *html.Node {
	var target *html.Node
	findElementByID := func(n *html.Node) bool {
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				target = n
				return false
			}
		}
		return true
	}
	post := func(n *html.Node) bool {
		return true
	}
	forEachNode(doc, findElementByID, post)
	return target
}

func summary(n *html.Node) string {
	if n == nil {
		return ""
	}
	var attributeText string
	for _, attr := range n.Attr {
		attributeText += " "
		attributeText += attr.Key + "=" + "'" + attr.Val + "'"
	}
	return fmt.Sprintf("<%s%s>", n.Data, attributeText)
}
