package treesort

import (
	"math/rand"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	tree := Sort(data)
	t.Log(tree)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}
