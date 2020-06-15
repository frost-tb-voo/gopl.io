// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package equal

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCyclecheck(t *testing.T) {
	one, oneAgain, two := 1, 1, 2

	type CyclePtr *CyclePtr
	var cyclePtr1, cyclePtr2 CyclePtr
	cyclePtr1 = &cyclePtr1
	cyclePtr2 = &cyclePtr2

	type CycleSlice []CycleSlice
	var cycleSlice = make(CycleSlice, 1)
	cycleSlice[0] = cycleSlice

	ch1, ch2 := make(chan int), make(chan int)
	var ch1ro <-chan int = ch1

	type mystring string

	var iface1, iface1Again, iface2 interface{} = &one, &oneAgain, &two

	for _, test := range []struct {
		x, y interface{}
		want bool
	}{
		// basic types
		{1, 1, false},
		{1, 2, false},   // different values
		{1, 1.0, false}, // different types
		{"foo", "foo", false},
		{"foo", "bar", false},
		{mystring("foo"), "foo", false}, // different types
		// slices
		{[]string{"foo"}, []string{"foo"}, false},
		{[]string{"foo"}, []string{"bar"}, false},
		{[]string{}, []string(nil), false},
		// slice cycles
		{cycleSlice, cycleSlice, true},
		// maps
		{
			map[string][]int{"foo": {1, 2, 3}},
			map[string][]int{"foo": {1, 2, 3}},
			false,
		},
		{
			map[string][]int{"foo": {1, 2, 3}},
			map[string][]int{"foo": {1, 2, 3, 4}},
			false,
		},
		{
			map[string][]int{},
			map[string][]int(nil),
			false,
		},
		// pointers
		{&one, &one, false},
		{&one, &two, false},
		{&one, &oneAgain, false},
		{new(bytes.Buffer), new(bytes.Buffer), false},
		// pointer cycles
		{cyclePtr1, cyclePtr1, true},
		{cyclePtr2, cyclePtr2, true},
		{cyclePtr1, cyclePtr2, true}, // they're deeply equal
		// functions
		{(func())(nil), (func())(nil), false},
		{(func())(nil), func() {}, false},
		{func() {}, func() {}, false},
		// arrays
		{[...]int{1, 2, 3}, [...]int{1, 2, 3}, false},
		{[...]int{1, 2, 3}, [...]int{1, 2, 4}, false},
		// channels
		{ch1, ch1, false},
		{ch1, ch2, false},
		{ch1ro, ch1, false}, // NOTE: not equal
		// interfaces
		{&iface1, &iface1, false},
		{&iface1, &iface2, false},
		{&iface1Again, &iface1, false},
	} {
		if Cyclecheck(test.x) != test.want {
			t.Errorf("Cyclecheck(%v, %v) = %t",
				test.x, test.y, !test.want)
		}
		if Cyclecheck(test.y) != test.want {
			t.Errorf("Cyclecheck(%v, %v) = %t",
				test.x, test.y, !test.want)
		}
	}
}

func Example_equal() {
	//!+
	fmt.Println(Cyclecheck([]int{1, 2, 3}))
	fmt.Println(Cyclecheck([]string{"foo"})) // "false"
	fmt.Println(Cyclecheck([]string(nil)))
	fmt.Println(Cyclecheck(map[string]int(nil)))
	//!-

	// Output:
	// false
	// false
	// false
	// false
}

func Example_equalCycle() {
	//!+cycle
	// Circular linked lists a -> b -> a and c -> c.
	type link struct {
		value string
		tail  *link
	}
	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c
	fmt.Println(Cyclecheck(a)) // "true"
	fmt.Println(Cyclecheck(b)) // "true"
	fmt.Println(Cyclecheck(c)) // "true"
	//!-cycle

	// Output:
	// true
	// true
	// true
}
