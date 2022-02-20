package solution

func sumOfThree(num int64) []int64 {
	if num%3 != 0 {
		return nil
	}
	m := num / 3
	return []int64{m - 1, m, m + 1}
}
