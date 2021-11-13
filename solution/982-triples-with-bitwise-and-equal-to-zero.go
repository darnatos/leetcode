package solution

func CountTriplets(nums []int) int {
	tuples := make(map[int]int, len(nums)*len(nums))
	for _, i := range nums {
		for j := range nums {
			tuples[i&nums[j]]++
		}
	}

	res := 0
	for k := range tuples {
		for _, i := range nums {
			if i&k == 0 {
				res += tuples[k]
			}
		}
	}

	return res
}
