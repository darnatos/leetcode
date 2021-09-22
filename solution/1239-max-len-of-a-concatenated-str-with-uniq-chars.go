package solution

func MaxLength(arr []string) int {
	m := make([]bool, 26)
	for i, s := range arr {
		for j := 0; j < len(s); j++ {
			if !m[s[j]-'a'] {
				m[s[j]-'a'] = true
			} else {
				arr[i] = ""
				break
			}
		}
		for k := 0; k < len(s); k++ {
			m[s[k]-'a'] = false
		}
	}
	cnt := 0
	mlDfs(arr, 0, m, &cnt)
	return cnt
}

func mlDfs(arr []string, j int, m []bool, cnt *int){
	*cnt = max(*cnt, mlCount(m))
	if j == len(arr) {
		return
	}
	for i := j; i < len(arr); i++ {
		s := arr[i]
		if mlCheck(s, m) {
			for k := range s {
				m[s[k]-'a'] = true
			}
			mlDfs(arr, i+1, m, cnt)
			for k := range s {
				m[s[k]-'a'] = false
			}
		}
	}
}

func mlCount(m []bool) int {
	cnt := 0
	for i := range m {
		if m[i] {
			cnt++
		}
	}
	return cnt
}

func mlCheck(s string, m []bool) bool {
	for i := range s {
		if m[s[i]-'a'] {
			return false
		}
	}
	return true
}
