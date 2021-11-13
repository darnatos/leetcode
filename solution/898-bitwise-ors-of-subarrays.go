package solution

func SubarrayBitwiseORs(arr []int) int {
	cur, cur2, res := make(map[int]struct{}, 30), make(map[int]struct{}), make(map[int]struct{})
	for _, i := range arr {
		cur2 = make(map[int]struct{}, 30)
		cur2[i] = struct{}{}
		for j := range cur {
			cur2[i|j] = struct{}{}
		}
		cur = cur2
		for j := range cur {
			res[j] = struct{}{}
		}
	}
	return len(res)
}
