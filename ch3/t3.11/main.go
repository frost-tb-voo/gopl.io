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
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func commaUpperThan1(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return commaUpperThan1(s[:n-3]) + "," + s[n-3:]
}

func commaLowerThan1(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return s[:3] + "," + commaLowerThan1(s[3:])
}

func comma(s string) string {
	var signC string
	var upperThan1 string
	var lowerThan1 string

	if s[0] == '+' || s[0] == '-' {
		signC = s[:1]
		s = s[1:]
	}
	upperThan1 = s

	var dotIndex int = strings.IndexRune(s, '.')
	if dotIndex > 0 {
		upperThan1 = s[:dotIndex]
		lowerThan1 = s[dotIndex+1:]
	}

	var result string = signC + commaUpperThan1(upperThan1)
	if len(lowerThan1) > 0 {
		result += "." + commaLowerThan1(lowerThan1)
	}
	return result
}

//!-
