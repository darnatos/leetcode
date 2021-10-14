package solution

func MajorityElement2(nums []int) []int {
	major1, count1, major2, count2 := 0, 0, 1, 0
	for _, v := range nums {
		if major1 == v {
			count1++
		} else if major2 == v {
			count2++
		} else if count1 == 0 {
			major1, count1 = v, 1
		} else if count2 == 0 {
			major2, count2 = v, 1
		} else {
			count1, count2 = count1-1, count2-1
		}
	}

	res := make([]int, 0, 2)
	c1, c2 := 0, 0
	n := len(nums) / 3
	for _, v := range nums {
		if v == major1 {
			c1++
		} else if v == major2 {
			c2++
		}
	}
	if c1 > n {
		res = append(res, major1)
	}
	if c2 > n {
		res = append(res, major2)
	}
	return res
}
