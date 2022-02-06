package solution

import "sort"

func smallestNumber(num int64) int64 {
	arr := make([]int64, 0, 16)
	var sign int64 = 1
	if num <= 0 {
		num = -num
		sign = -1
	}
	for num > 0 {
		arr = append(arr, num%10)
		num /= 10
	}
	var res int64
	if sign < 0 {
		sort.Slice(arr, func(i, j int) bool { return arr[i] > arr[j] })
		res = 0
		for i := range arr {
			res = 10*res + arr[i]
		}
	} else {
		sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
		res = 0
		i := 0
		for i < len(arr) && arr[i] == 0 {
			i++
		}
		res = arr[i]
		for j := 0; j < i; j++ {
			res *= 10
		}
		i++

		for i < len(arr) {
			res = 10*res + arr[i]
			i++
		}
	}
	return sign * res
}
