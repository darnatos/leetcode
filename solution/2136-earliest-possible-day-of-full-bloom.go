package solution

import "sort"

func EarliestFullBloom(plantTime []int, growTime []int) int {
	n := len(plantTime)
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = []int{growTime[i], plantTime[i]}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] > arr[j][0]
	})
	plant := 0
	res := 0
	for i := range arr {
		plant += arr[i][1]
		res = max(res, plant+arr[i][0])
	}

	return res
}
