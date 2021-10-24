package solution

import "strconv"

func NextBeautifulNumber(n int) int {
	for i := n + 1; i <= 1224444; i++ {
		if isBalanced(i) {
			return i
		}
	}
	return 1224444
}

func isBalanced(i int) bool {
	str := strconv.Itoa(i)
	cc := make(map[byte]byte)
	for i := range str {
		cc[str[i]-'0']++
	}
	for i := range cc {
		if i != cc[i] {
			return false
		}
	}
	return true
}
