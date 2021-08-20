package util

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Abs(s int) int {
	if s < 0 {
		return -s
	}
	return s
}
