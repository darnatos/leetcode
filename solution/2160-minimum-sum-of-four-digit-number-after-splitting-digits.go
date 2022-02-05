package solution

import "sort"

func minimumSum(num int) int {
	m := make([]int, 0, 4)
	for num > 0 {
		m = append(m, num%10)
		num /= 10
	}
	sort.Ints(m)

	return 10*(m[0]+m[1]) + m[2] + m[3]
}
