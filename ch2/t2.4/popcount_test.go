package popcount_test

import (
	"testing"

	popcount "github.com/frost-tb-voo/gopl.io/ch2/t2.4"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}
