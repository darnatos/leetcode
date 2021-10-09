package solution

import "sort"

func CanReorderDoubled(arr []int) bool {
	c := make(map[int]int)
	for i := range arr {
		c[arr[i]]++
	}
	keys := arr[:0]
	for i := range c {
		keys = append(keys, i)
	}

	sort.Slice(keys, func(i, j int) bool { return abs(keys[i]) < abs(keys[j]) })

	for _, i := range keys {
		if c[i] > c[i*2] {
			return false
		}
		c[i*2] -= c[i]
	}
	return true
}
