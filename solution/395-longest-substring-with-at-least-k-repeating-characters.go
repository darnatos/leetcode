package solution

func longestSubstring(s string, k int) int {
	res := 0
	i, j, uniqueNumber, noLessThanK := 0, 0, 0, 0
	c := make([]int, 26)
	for h := 1; h <= 26; h++ {
		i, j, uniqueNumber, noLessThanK = 0, 0, 0, 0
		for a := range c {
			c[a] = 0
		}

		for j < len(s) {
			if uniqueNumber <= h {
				if c[s[j]-'a'] == 0 {
					uniqueNumber++
				}
				c[s[j]-'a']++
				if c[s[j]-'a'] == k {
					noLessThanK++
				}
				j++
			} else {
				if c[s[i]-'a'] == k {
					noLessThanK--
				}
				c[s[i]-'a']--
				if c[s[i]-'a'] == 0 {
					uniqueNumber--
				}
				i++
			}
			if h == uniqueNumber && uniqueNumber == noLessThanK {
				res = max(res, j-i)
			}
		}
	}

	return res
}
