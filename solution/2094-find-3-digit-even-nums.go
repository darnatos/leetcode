package solution

import "sort"

func FindEvenNumbers(digits []int) []int {
	s := make(map[int]bool)
	for i := 0; i < len(digits); i++ {
		if digits[i] == 0 {
			continue
		}
		for j := 0; j < len(digits); j++ {
			if i == j {
				continue
			}
			for k := 0; k < len(digits); k++ {
				if j == k || i == k {
					continue
				}
				tmp := digits[i]*100 + digits[j]*10 + digits[k]
				if tmp&1 == 0 {
					s[tmp] = true
				}
			}
		}
	}
	res := make([]int, 0, len(s))
	for i := range s {
		res = append(res, i)
	}
	sort.Ints(res)
	return res
}
