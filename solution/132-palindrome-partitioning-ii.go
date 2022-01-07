package solution

func minCut(s string) int {
	n := len(s)
	cuts := make([]int, n+1)
	for i := range cuts {
		cuts[i] = i - 1
	}

	for i := 0; i < len(s); i++ {
		for j := 0; i-j >= 0 && i+j < n && s[i-j] == s[i+j]; j++ {
			cuts[i+j+1] = min(cuts[i+j+1], cuts[i-j]+1)
		}
		for j := 1; i-j+1 >= 0 && i+j < n && s[i-j+1] == s[i+j]; j++ {
			cuts[i+j+1] = min(cuts[i+j+1], cuts[i-j+1]+1)
		}
	}

	return cuts[n]
}
