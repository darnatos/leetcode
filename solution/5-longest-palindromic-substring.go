package solution

import "fmt"

func LongestPalindrome(s string) string {
	t := []byte{'$', '#'}
	for i := range s {
		t = append(t, s[i], '#')
	}
	t = append(t, '^')

	p := make([]int, len(t)-1)
	var id, mx, resId, resMx int
	for i := 1; i < len(t)-1; i++ {
		if mx > i {
			p[i] = min(p[id*2-i], mx-i)
		} else {
			p[i] = 1
		}

		for t[i+p[i]] == t[i-p[i]] {
			p[i]++
		}

		//fmt.Println(string(t))
		//fmt.Println(p)
		//fmt.Println(s[(resId-resMx)/2 : (resId+resMx-1)/2])
		fmt.Println(id, mx, resId, resMx)
		if mx < i+p[i] {
			mx = i + p[i]
			id = i
		}
		if resMx < p[i] {
			resMx = p[i]
			resId = i
		}
	}

	return s[(resId-resMx)/2 : (resId+resMx-1)/2]
}
