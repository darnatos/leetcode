package solution

func PossiblyEquals(s1 string, s2 string) bool {
	visited := [41][41][2000]bool{}

	return peDfs(s1, s2, 0, 0, 0, &visited)
}

func peDfs(s1, s2 string, i, j, diff int, visited *[41][41][2000]bool) bool {
	var processDigits func(s string, p *int, sign int) bool
	processDigits = func(s string, p *int, sign int) bool {
		for val := 0; *p < len(s) && isDigit(s[*p]); {
			val = 10*val + int(s[*p]-'0')
			*p++
			if peDfs(s1, s2, i, j, diff+sign*val, visited) {
				return true
			}
		}
		return false
	}

	if i == len(s1) && j == len(s2) {
		return diff == 0
	}
	if !visited[i][j][1000+diff] {
		visited[i][j][1000+diff] = true
		if i < len(s1) && isDigit(s1[i]) {
			return processDigits(s1, &i, -1)
		}
		if j < len(s2) && isDigit(s2[j]) {
			return processDigits(s2, &j, 1)
		}
		if diff > 0 {
			return i < len(s1) && peDfs(s1, s2, i+1, j, diff-1, visited)
		}
		if diff < 0 {
			return j < len(s2) && peDfs(s1, s2, i, j+1, diff+1, visited)
		}
		return i < len(s1) && j < len(s2) && s1[i] == s2[j] && peDfs(s1, s2, i+1, j+1, diff, visited)
	}
	return false
}
