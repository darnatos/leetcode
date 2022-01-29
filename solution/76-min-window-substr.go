package solution

import "math"

func MinWindow(s string, t string) string {
	m, n := len(s), len(t)
	if m < n {
		return ""
	}

	tm := [128]int{}
	for i := range t {
		tm[t[i]]++
	}
	head, d := 0, math.MaxInt32
	for i, j := 0, 0; j < m; {
		if tm[s[j]] > 0 {
			n--
		}
		tm[s[j]]--
		j++
		for n == 0 {
			if j-i < d {
				head = i
				d = j - i
			}
			if tm[s[i]] == 0 {
				n++
			}
			tm[s[i]]++
			i++
		}
	}
	if d == math.MaxInt32 {
		return ""
	}
	return string(s[head : head+d])
}
