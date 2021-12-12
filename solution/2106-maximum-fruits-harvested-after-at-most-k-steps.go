package solution

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	sum, res := 0, 0
	l := 0
	for l < len(fruits) && startPos-k > fruits[l][0] {
		l++
	}
	for r := l; r < len(fruits) && fruits[r][0] <= startPos+k; r++ {
		sum += fruits[r][1]

		for startPos-fruits[l][0]*2+fruits[r][0] > k && fruits[r][0]*2-startPos-fruits[l][0] > k {
			sum -= fruits[l][1]
			l++
		}

		res = max(res, sum)
	}

	return res
}
