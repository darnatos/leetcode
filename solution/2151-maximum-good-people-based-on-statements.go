package solution

import "math/bits"

func maximumGood(statements [][]int) int {
	n := len(statements)
	all := 1 << n
	res := 0
	for i := 1; i < all; i++ {
		cnt := bits.OnesCount(uint(i))
		if res < cnt && isValid(statements, i) {
			res = cnt
		}
	}
	return res
}

func isValid(statements [][]int, mask int) bool {
	for i := range statements {
		if (1<<i)&mask > 0 {
			for j := range statements {
				if statements[i][j] == 2 {
					continue
				}
				if statements[i][j] == 1 && (1<<j)&mask == 0 || statements[i][j] == 0 && (1<<j)&mask > 0 {
					return false
				}
			}
		}
	}
	return true
}
