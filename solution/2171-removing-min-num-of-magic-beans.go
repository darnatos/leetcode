package solution

import "sort"

func minimumRemoval(beans []int) int64 {
	n := len(beans)
	sort.Ints(beans)
	sum := int64(0)
	for i := range beans {
		sum += int64(beans[i])
	}
	res := sum
	pre := 0
	for i := 0; i < n; i++ {
		sum += int64(pre - (beans[i]-pre)*(n-i))
		pre = beans[i]
		res = min64(res, sum)
	}
	return res
}
