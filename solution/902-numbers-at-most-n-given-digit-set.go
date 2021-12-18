package solution

import (
	"math"
	"strconv"
)

func atMostNGivenDigitSet(digits []string, n int) int {
	ns := strconv.Itoa(n)
	dSize := len(digits)
	res := 0
	for i := 1; i < len(ns); i++ {
		res += int(math.Pow(float64(dSize), float64(i)))
	}

	for i := 0; i < len(ns); i++ {
		hasSame := false
		for _, d := range digits {
			if d[0] < ns[i] {
				res += int(math.Pow(float64(dSize), float64(len(ns)-i-1)))
			} else if d[0] == ns[i] {
				hasSame = true
			} else {
				// `digits` is sorted in non-decreasing order
				break
			}
		}
		if !hasSame {
			return res
		}
	}

	return res + 1
}
