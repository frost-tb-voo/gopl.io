// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var rr1 []rune = []rune(s)
	var bb bytes.Buffer
	var count int = 0
	for ii := len(rr1) - 1; ii >= 0; ii-- {
		if count > 0 && count%3 == 0 {
			bb.WriteRune(',')
		}
		bb.WriteRune(rr1[ii])
		count++
	}
	var rr2 []rune = []rune(bb.String())
	for i, j := 0, len(rr2)-1; i < j; i, j = i+1, j-1 {
		rr2[i], rr2[j] = rr2[j], rr2[i]
	}
	return string(rr2)
}

//!-
