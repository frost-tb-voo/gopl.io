package main

import (
	"testing"
)

func TestCountWordsAndImages(tt *testing.T) {
	words, images, err := CountWordsAndImages("https://www.google.com/")
	if err != nil {
		tt.Fatalf("%v\n", err)
	}
	tt.Logf("words:\t%v\n", words)
	tt.Logf("images:\t%v\n", images)
}
