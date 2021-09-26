package solution

import "sort"

func FindLeastNumOfUniqueInts(arr []int, k int) int {
	c := make(map[int]int)
	for i := range arr {
		c[arr[i]]++
	}
	q := make([][2]int, 0, len(c))

	for n, v := range c {
		q = append(q, [2]int{n, v})
	}
	sort.Slice(q, func(a, b int) bool { return q[a][1] < q[b][1] })

	res := len(q)
	for i := range q {
		if q[i][1] <= k {
			k -= q[i][1]
			res--
		} else {
			break
		}
	}

	return res
}
