package solution

func repeatLimitedString(s string, repeatLimit int) string {
	cnt := [26]int{}
	for i := range s {
		cnt[s[i]-'a']++
	}
	n := len(s)
	res := make([]byte, 0)
	pre := 26
	var cur, next, rpt int
	for n > 0 {
		cur, next, rpt = 26, 26, 0
		for j := 25; j >= 0; j-- {
			if cnt[j] > 0 && pre != j {
				cur = j
				if cnt[j] >= repeatLimit {
					rpt = repeatLimit
					cnt[j] -= repeatLimit
				} else {
					rpt = cnt[j]
					cnt[j] = 0
				}
				n -= rpt
				if cnt[j] > 0 {
					for j--; j >= 0; j-- {
						if cnt[j] > 0 {
							next = j
							cnt[j]--
							n--
							break
						}
					}
				}
				break
			}
		}
		if cur >= 26 {
			break
		}
		for i := 0; i < rpt; i++ {
			res = append(res, byte(cur+'a'))
		}
		pre = cur
		if next >= 26 {
			continue
		}
		res = append(res, byte(next+'a'))
		pre = next
	}
	return string(res)
}
