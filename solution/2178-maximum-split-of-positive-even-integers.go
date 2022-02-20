package solution

func maximumEvenSplit(finalSum int64) []int64 {
	if finalSum&1 == 1 {
		return nil
	}
	cur := int64(2)
	res := make([]int64, 0)
	for finalSum >= cur {
		res = append(res, cur)
		finalSum -= cur
		cur += 2
	}
	res[len(res)-1] += finalSum
	return res
}
