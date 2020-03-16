// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
//!+
package popcount

import (
	"sync"
)

type countSync struct {
	sync.RWMutex
	init bool
	pc   [256]byte // pc[i] is the population count of i.
}

var count = countSync{init: false}

func init() {
	count.RLock()
	if count.init {
		count.RUnlock()
		return
	}
	count.RUnlock()
	count.Lock()
	defer count.Unlock()
	pc := count.pc
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	count.RLock()
	defer count.RUnlock()
	pc := count.pc
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

//!-
