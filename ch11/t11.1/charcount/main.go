// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	charcount(os.Stdin, os.Stdout)
}

func charcount(stdin io.Reader, stdout io.Writer) {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Fprintf(stdout, "rune\tcount\n")
	for c, n := range counts {
		fmt.Fprintf(stdout, "%q\t%d\n", c, n)
	}
	fmt.Fprintf(stdout, "\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Fprintf(stdout, "%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Fprintf(stdout, "\n%d invalid UTF-8 characters\n", invalid)
	}
}

//!-
