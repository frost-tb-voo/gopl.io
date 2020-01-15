package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	words, images, err := CountWordsAndImages("https://www.google.com/")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "words:\t%v\n", words)
	fmt.Fprintf(os.Stdout, "images:\t%v\n", images)
}

// CountWordsAndImages is ..
// HTML document に対する HTML GET リクエストを url へ行いそのドキュメント内に含まれる単語と画像の数を返す
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.TextNode {
		input := bufio.NewScanner(strings.NewReader(n.Data))
		input.Split(bufio.ScanWords)
		for input.Scan() {
			// text := input.Text()
			words++
		}
	}
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		cWords, cImages := countWordsAndImages(c)
		words += cWords
		images += cImages
	}
	return
}
