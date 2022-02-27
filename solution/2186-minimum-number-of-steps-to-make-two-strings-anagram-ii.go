package solution

func minSteps(s string, t string) int {
	cnt := [26]int{}
	for i := range s {
		cnt[s[i]-'a']++
	}
	for j := range t {
		cnt[t[j]-'a']--
	}
	res := 0
	for i := range cnt {
		if cnt[i] >= 0 {
			res += cnt[i]
		} else {
			res -= cnt[i]
		}
	}
	return res
}
