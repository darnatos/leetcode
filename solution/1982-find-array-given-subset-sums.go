package solution

import "sort"

func RecoverArray(n int, sums []int) []int {
	sort.Ints(sums)

	freq := make(map[int]int, len(sums)/2)
	ss0, ss1 := make([]int, len(sums)/2), make([]int, len(sums)/2)
	res := make([]int, n)

	for n = n - 1; n >= 0; n-- {
		ss0, ss1 = ss0[:0], ss1[:0]
		diff := sums[1] - sums[0]
		on := false

		for _, x := range sums {
			if freq[x] == 0 {
				ss0 = append(ss0, x)
				freq[x+diff]++
				if x == 0 {
					on = true
				}
			} else {
				ss1 = append(ss1, x)
				freq[x]--
			}
		}

		if on {
			sums = ss0
			res[n] = diff
		} else {
			sums = ss1
			res[n] = -diff
		}
		// fmt.Println(res, diff, freq, on, ss0, ss1)
	}

	return res
}
