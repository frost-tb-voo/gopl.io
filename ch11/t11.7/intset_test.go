package intset_test

import (
	"fmt"
	"testing"

	"github.com/frost-tb-voo/gopl.io/ch6/intset"
	"github.com/golang-collections/collections/set"
)

func BenchmarkMapsetAdd(bb *testing.B) {
	for ii := 0; ii < bb.N; ii++ {
		x := set.New()
		x.Insert(1)
		x.Insert(144)
		x.Insert(9)
	}
}

func BenchmarkIntsetAdd(bb *testing.B) {
	for ii := 0; ii < bb.N; ii++ {
		var x intset.IntSet
		x.Add(1)
		x.Add(144)
		x.Add(9)
	}
}

func BenchmarkMapsetString(bb *testing.B) {
	x := set.New()
	x.Insert(1)
	x.Insert(144)
	x.Insert(9)
	for ii := 0; ii < bb.N; ii++ {
		fmt.Sprintf("%v", x)
	}
}

func BenchmarkIntsetString(bb *testing.B) {
	var x intset.IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	for ii := 0; ii < bb.N; ii++ {
		x.String()
	}
}

func BenchmarkMapsetUnionWith(bb *testing.B) {
	x := set.New()
	x.Insert(1)
	x.Insert(144)
	x.Insert(9)

	y := set.New()
	y.Insert(9)
	y.Insert(42)
	for ii := 0; ii < bb.N; ii++ {
		x.Union(y)
	}
}

func BenchmarkIntsetUnionWith(bb *testing.B) {
	var x, y intset.IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)

	y.Add(9)
	y.Add(42)
	for ii := 0; ii < bb.N; ii++ {
		x.UnionWith(&y)
	}
}

func BenchmarkMapsetHas(bb *testing.B) {
	x := set.New()
	x.Insert(1)
	x.Insert(144)
	x.Insert(9)

	y := set.New()
	y.Insert(9)
	y.Insert(42)

	x = x.Union(y)
	for ii := 0; ii < bb.N; ii++ {
		x.Has(9)
		x.Has(123)
	}
}

func BenchmarkIntsetHas(bb *testing.B) {
	var x, y intset.IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)

	y.Add(9)
	y.Add(42)

	x.UnionWith(&y)
	for ii := 0; ii < bb.N; ii++ {
		x.Has(9)
		x.Has(123)
	}
}

// goos: linux
// goarch: amd64
// BenchmarkMapsetAdd-4         	 3649704	       295 ns/op	     200 B/op	       3 allocs/op
// BenchmarkIntsetAdd-4         	 9197610	       125 ns/op	      56 B/op	       3 allocs/op
// BenchmarkMapsetString-4      	  768840	      1576 ns/op	     416 B/op	       9 allocs/op
// BenchmarkIntsetString-4      	 1976283	       609 ns/op	     152 B/op	       6 allocs/op
// BenchmarkMapsetUnionWith-4   	 2249325	       530 ns/op	     200 B/op	       3 allocs/op
// BenchmarkIntsetUnionWith-4   	265160625	         4.36 ns/op	       0 B/op	       0 allocs/op
// BenchmarkMapsetHas-4         	24126332	        49.6 ns/op	       0 B/op	       0 allocs/op
// BenchmarkIntsetHas-4         	1000000000	         0.293 ns/op	       0 B/op	       0 allocs/op
