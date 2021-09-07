package solution

func TopKFrequent(nums []int, k int) []int {
	f := make(map[int]int, k)
	for _, v := range nums {
		f[v]++
	}

	buckets := make([][]int, len(nums)+1)
	for i, v := range f {
		if buckets[v] == nil {
			buckets[v] = make([]int, 0, 2)
		}
		buckets[v] = append(buckets[v], i)
	}

	res := make([]int, 0, k)
	for i := len(nums); len(res) < k; i-- {
		if buckets[i] != nil {
			res = append(res, buckets[i]...)
		}
	}

	return res[:k]
}
