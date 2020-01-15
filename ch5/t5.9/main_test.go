package main

import (
	"testing"
)

func TestVisit(tt *testing.T) {
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
	tt.Logf("%v --> \n", ss)
	text = expand(ss, ff1)
	tt.Logf("%v\n", text)
	text = expand(ss, ff2)
	tt.Logf("%v\n", text)
}
