package solution

func minimumOperations(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	odd, even := make(map[int]int), make(map[int]int)
	for i := range nums {
		if i&1 == 1 {
			odd[nums[i]]++
		} else {
			even[nums[i]]++
		}
	}

	fo, so := 0, 0
	for i := range odd {
		if odd[i] > odd[fo] {
			so = fo
			fo = i
		} else if odd[i] > odd[so] {
			so = i
		}
	}

	fe, se := 0, 0
	for i := range even {
		if even[i] > even[fe] {
			se = fe
			fe = i
		} else if even[i] > even[se] {
			se = i
		}
	}
	if fo == fe {
		return len(nums) - max(odd[fo]+even[se], odd[so]+even[fe])
	}
	return len(nums) - (odd[fo] + even[fe])
}
