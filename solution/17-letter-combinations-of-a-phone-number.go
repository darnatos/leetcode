package solution

func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	dp := []string{""}
	for i := range digits {
		dpp := make([]string, 0, len(dp)*4)
		for _, ch := range toLet(digits[i]) {
			for j := range dp {
				dpp = append(dpp, dp[j]+string(ch))
			}
		}
		dp = dpp
	}
	return dp
}

func toLet(d byte) []byte {
	switch d {
	case '9':
		return []byte{'w', 'x', 'y', 'z'}
	case '8':
		return []byte{'t', 'u', 'v'}
	case '7':
		return []byte{'p', 'q', 'r', 's'}
	default:
		res := make([]byte, 0, 4)
		ch := (d-'2')*3 + 'a'
		for i := ch; i < ch+3; i++ {
			res = append(res, i)
		}
		return res
	}
}
