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
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	forEachNode(doc, startElement, endElement)
	//!-call

	return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

//!-forEachNode

//!+startend
var depth int

// void elements : area, base, br, col, embed, hr, img, input, link, meta, param, source, track, wbr
// https://html.spec.whatwg.org/multipage/syntax.html#void-elements
var voidElements []string = []string{"area", "base", "br", "col", "embed", "hr", "img", "input", "link", "meta", "param", "source", "track", "wbr"}

func isVoid(tagName string) bool {
	for _, element := range voidElements {
		if element == tagName {
			return true
		}
	}
	return false
}

func startElement(n *html.Node) {
	var attributeText string
	for _, attr := range n.Attr {
		attributeText += " "
		attributeText += attr.Key + "=" + "'" + attr.Val + "'"
	}
	if n.Type == html.ElementNode {
		if isVoid(n.Data) {
			if n.FirstChild != nil {
				panic(n.Data + " is void element but has childs")
			}
			fmt.Printf("%*s<%s%s/>\n", depth*2, "", n.Data, attributeText)
		} else {
			fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, attributeText)
			depth++
		}
	}
	if n.Type == html.ErrorNode || n.Type == html.CommentNode {
		fmt.Printf("%*s<!-- %s -->\n", depth*2, "", n.Data)
	}
	if n.Type == html.TextNode {
		fmt.Printf("%*s%s\n", depth*2, "", n.Data)
	}
	if n.Type == html.DocumentNode {
	}
	if n.Type == html.DoctypeNode {
		fmt.Printf("%*s<!%s%s>\n", depth*2, "", n.Data, attributeText)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		if isVoid(n.Data) {
		} else {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}

//!-startend
