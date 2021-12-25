package solution

import (
	"fmt"
	"strconv"
)

func AbbreviateProduct(left int, right int) string {
	all, pre, suf := 1, 1, 1
	flag := true
	zeros := 0

	for i := left; i <= right; i++ {
		if flag {
			all *= i
			for all%10 == 0 {
				all /= 10
			}
			if all >= int(1e10) {
				flag = false
			}
		}

		pre *= i
		suf *= i

		for pre >= int(1e12) {
			pre /= 10
		}
		for suf%10 == 0 {
			suf /= 10
			zeros++
		}
		if suf >= int(1e10) {
			suf %= int(1e10)
		}
	}
	for pre >= 100000 {
		pre /= 10
	}
	if flag {
		return strconv.Itoa(all) + "e" + strconv.Itoa(zeros)
	}
	return strconv.Itoa(pre) + "..." + fmt.Sprintf("%05d", suf%100000) + "e" + strconv.Itoa(zeros)
}
