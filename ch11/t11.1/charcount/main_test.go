package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestCharcount(tt *testing.T) {
	var tests = []struct {
		in   string
		want string
	}{
		{"aiueeoao",
			`rune	count
'a'	2
'i'	1
'u'	1
'e'	2
'o'	2

len	count
1	8
2	0
3	0
4	0
`},
	}
	for _, test := range tests {
		descr := fmt.Sprintf("charcount(%q)", test.in)
		in := strings.NewReader(test.in)
		out := new(bytes.Buffer)
		charcount(in, out)
		got := out.String()
		if got != test.want {
			tt.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
