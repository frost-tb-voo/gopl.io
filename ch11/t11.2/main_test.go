package intset_test

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"gopl.io/ch6/intset"
)

func TestAdd(tt *testing.T) {
	var tests = []struct {
		in []int
	}{
		{
			[]int{1, 144, 9, 42},
		},
	}

	for _, test := range tests {
		var x intset.IntSet
		var wantX map[int]bool = map[int]bool{}
		for _, num := range test.in {
			x.Add(num)
			wantX[num] = true
		}

		var want []int
		for key, _ := range wantX {
			want = append(want, key)
		}
		sort.SliceStable(want, func(i, j int) bool { return want[i] < want[j] })

		wantStr := fmt.Sprintf("%v", want)
		wantStr = strings.ReplaceAll(wantStr, "[", "{")
		wantStr = strings.ReplaceAll(wantStr, "]", "}")
		gotStr := x.String()
		if gotStr != wantStr {
			tt.Errorf("want:%v, but got:%v", wantStr, gotStr)
		}
	}
}

func TestUnionWith(tt *testing.T) {
	var tests = []struct {
		arg1 []int
		arg2 []int
	}{
		{
			[]int{1, 144, 9, 42},
			[]int{1, 13, 91, 42},
		},
	}

	for _, test := range tests {
		var x intset.IntSet
		var y intset.IntSet
		var wantX map[int]bool = map[int]bool{}
		var wantY map[int]bool = map[int]bool{}
		for _, num := range test.arg1 {
			x.Add(num)
			wantX[num] = true
		}
		for _, num := range test.arg2 {
			y.Add(num)
			wantY[num] = true
		}

		for key, _ := range wantY {
			wantX[key] = true
		}
		x.UnionWith(&y)

		var want []int
		for key, _ := range wantX {
			want = append(want, key)
		}
		sort.SliceStable(want, func(i, j int) bool { return want[i] < want[j] })

		wantStr := fmt.Sprintf("%v", want)
		wantStr = strings.ReplaceAll(wantStr, "[", "{")
		wantStr = strings.ReplaceAll(wantStr, "]", "}")
		gotStr := x.String()
		if gotStr != wantStr {
			tt.Errorf("want:%v, but got:%v", wantStr, gotStr)
		}
	}
}
