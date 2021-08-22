package solution

import "sort"

func ArrayRankTransform(arr []int) []int {
	a := make([]int, len(arr))

	set := make(map[int]bool, len(arr))

	idx := 0
	for _, v := range arr {
		if set[v] != true {
			set[v] = true
			a[idx] = v
			idx++
		}
	}
	a = a[:idx]

	sort.Ints(a)

	valueToRank := make(map[int]int, len(a))
	rank := 1
	for i := range a {
		if _, ok := valueToRank[a[i]]; !ok {
			valueToRank[a[i]] = rank
			rank++
		}
	}

	for i := range arr {
		arr[i] = valueToRank[arr[i]]
	}

	return arr
}
