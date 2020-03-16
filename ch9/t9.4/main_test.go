package main

import (
	"math"
	"testing"
)

func Test(tt *testing.T) {
	for ii := 0; ii < 6; ii += 1 {
		length := int(math.Pow10(ii))
		chain(length)
	}
	pow5 := int(math.Pow10(5))
	for ii := pow5; ii < pow5*10; ii += pow5 {
		length := ii
		chain(length)
	}
}
