package popcount

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	xx := x
	count := 0
	for ii := uint(0); ii < 64; ii++ {
		count += int(xx & 1)
		xx = xx >> 1
	}
	return count
}
