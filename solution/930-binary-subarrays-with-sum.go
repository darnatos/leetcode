package solution

func NumSubarraysWithSum2P(nums []int, goal int) int {
	var i0, sum, offset, res int

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum == goal {
			for nums[i0] == 0 && i0 < i {
				i0++
				offset++
			}
			res += offset + 1
		} else if sum > goal {
			for sum > goal {
				sum -= nums[i0]
				i0++
			}
			offset = 0
			if i0 > i {
				continue
			}
			sum -= nums[i]
			i--
		}
	}

	return res
}

func NumSubarraysWithSumMap(nums []int, goal int) int {
	m := make(map[int]int)
	m[0] = 1
	sum := 0
	res := 0

	for i := range nums {
		sum += nums[i]
		res += m[sum-goal]
		m[sum]++
	}

	return res
}
