package popcount_test

import (
	"testing"

	bitshift "github.com/frost-tb-voo/gopl.io/ch2/t2.4"
	bitclear "github.com/frost-tb-voo/gopl.io/ch2/t2.5"
	popcount "gopl.io/ch2/popcount"
)

var pc [256]byte

func popcountInit() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func BitCount(x uint64) int {
	// Hacker's Delight, Figure 5-2.
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	return int(x & 0x7f)
}

// -- Benchmarks --

func BenchmarkPopCountInit(b *testing.B) {
	for ii := 0; ii < b.N; ii++ {
		popcountInit()
	}
}

func BenchmarkPopCount(b *testing.B) {
	for ii := 0; ii < b.N; ii++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkBitCount(b *testing.B) {
	for ii := 0; ii < b.N; ii++ {
		BitCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByClearing(b *testing.B) {
	for ii := 0; ii < b.N; ii++ {
		bitclear.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByShifting(b *testing.B) {
	for ii := 0; ii < b.N; ii++ {
		bitshift.PopCount(0x1234567890ABCDEF)
	}
}

// goos: linux
// goarch: amd64
// BenchmarkPopCountInit
// BenchmarkPopCountInit-4         	 3145381	       378 ns/op	       0 B/op	       0 allocs/op
// BenchmarkPopCount
// BenchmarkPopCount-4             	1000000000	         0.283 ns/op	       0 B/op	       0 allocs/op
// BenchmarkBitCount
// BenchmarkBitCount-4             	1000000000	         0.285 ns/op	       0 B/op	       0 allocs/op
// BenchmarkPopCountByClearing
// BenchmarkPopCountByClearing-4   	49692204	        22.9 ns/op	       0 B/op	       0 allocs/op
// BenchmarkPopCountByShifting
// BenchmarkPopCountByShifting-4   	20225055	        57.8 ns/op	       0 B/op	       0 allocs/op

// 378/22.9=16.5..
// 378/57.8=6.5..
