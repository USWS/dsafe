package dsafe

import "math/rand"

// [min,max)
func RangeRand(min int, max int) int {
	r := rand.Intn(max)
	if r < min {
		return r + min
	}
	return r
}
