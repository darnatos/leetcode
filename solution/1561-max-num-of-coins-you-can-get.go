package solution

import "sort"

func MaxCoins(piles []int) int {
	sort.Ints(piles)

	res := 0
	for i := len(piles) / 3; i <= len(piles)-2; i += 2 {
		res += piles[i]
	}

	return res
}
