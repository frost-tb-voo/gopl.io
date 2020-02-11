package treesort

import (
	"strconv"
)

//!+
type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) *tree {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
	return root
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree) String() string {
	return t.stringInternal([]string{})
}

func (t *tree) stringInternal(depthes []string) string {
	var ss string
	if t.value != 0 {
		for _, depth := range(depthes) {
			ss += depth
		}
		ss += " " + strconv.Itoa(t.value) + "\n"
	}
	if t.left != nil {
		depthes = append(depthes, "l")
		ss += t.left.stringInternal(depthes)
	}
	if t.right != nil {
		if t.left != nil {
			depthes[len(depthes)-1] = "r"
		} else {
			depthes = append(depthes, "r")
		}
		ss += t.right.stringInternal(depthes)
	}
	return ss
}

//!-
