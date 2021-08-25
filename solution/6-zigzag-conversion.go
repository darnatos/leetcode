package solution

func ZigzagConvert(s string, numRows int) string {
	if numRows == 1 || len(s) == 1 {
		return s
	}
	flag := true
	v := make([][]byte, numRows)
	for i := range v {
		v[i] = make([]byte, 0, len(s)/numRows)
	}
	for i, j := 0, -1; i < len(s); i++ {
		if flag {
			j++
		} else {
			j--
		}
		v[j] = append(v[j], s[i])
		if j == numRows-1 {
			flag = false
		} else if j == 0 {
			flag = true
		}
	}
	s = ""
	for i := range v {
		s += string(v[i])
	}
	return s
}
