package solution

import "strconv"

func CountAndSay(n int) string {
	prev := []rune{'1'}
	for n > 1 {
		c, j, tmp := 1, 1, prev[0]
		cur := make([]rune, 0, len(prev)*2)
		for ;j < len(prev); j++ {
			if prev[j] != tmp {
				cur = append(cur, []rune(strconv.Itoa(c))...)
				cur = append(cur, tmp)
				tmp = prev[j]
				c = 1
			} else {
				c++
			}
		}
		cur = append(cur, []rune(strconv.Itoa(c))...)
		cur = append(cur, tmp)

		prev = cur
		n--
	}
	return string(prev)
}
