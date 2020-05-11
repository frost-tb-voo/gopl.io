package popcount

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	n := 0
	for x != 0 {
		x = x & (x - 1)
		n++
	}
	return n
}
