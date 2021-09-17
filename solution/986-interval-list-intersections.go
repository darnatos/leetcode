package solution

func IntervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	res := make([][]int, 0, 4)
	i, j, e, ee := 0, 0, -1, -1
	for i < len(firstList) && j < len(secondList) {
		a1, a2, b1, b2 := firstList[i][0], firstList[i][1], secondList[j][0], secondList[j][1]
		ss, s := minPair(a1, b1)
		if ee == ss {
			res = append(res, []int{ee, ee})
		}
		if a2 < b1 {
			i++
		} else if b2 < a1 {
			j++
		} else {
			e, ee = minPair(a2, b2)
			res = append(res, []int{s, e})
			if a2 <= b2 {
				i++
			}
			if b2 <= a2 {
				j++
			}
		}
	}
	return res
}

func minPair(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}