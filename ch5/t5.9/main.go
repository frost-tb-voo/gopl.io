package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	ss := "$foo played $hoge games."
	ff1 := func(key string) string {
		if key == "foo" {
			return "I"
		}
		if key == "hoge" {
			return "99"
		}
		return ""
	}
	ff2 := func(key string) string {
		if key == "foo" {
			return "We"
		}
		if key == "hoge" {
			return "hard"
		}
		return ""
	}
	var text string
	fmt.Fprintf(os.Stdout, "%v --> \n", ss)
	text = expand(ss, ff1)
	fmt.Fprintf(os.Stdout, "%v\n", text)
	text = expand(ss, ff2)
	fmt.Fprintf(os.Stdout, "%v\n", text)
}

// expand is ..
// 文字列 s 内のそれぞれの部分文字列 "$foo" を f("foo") が返すテキストで置換
func expand(s string, f func(string) string) string {
	var expanded string
	input := bufio.NewScanner(strings.NewReader(s))
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		if len(expanded) > 0 {
			expanded += " "
		}
		if strings.HasPrefix(word, "$") {
			word = f(strings.TrimPrefix(word, "$"))
		}
		expanded += word
	}
	return expanded
}
