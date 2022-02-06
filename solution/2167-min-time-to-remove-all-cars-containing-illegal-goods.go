package solution

func minimumTime(s string) int {
	n := len(s)
	// res can't be larger than length of s (remove all cars from right to left)
	res := n
	leftMin := 0
	for i := 0; i < n; i++ {
		// remove illegal car from anywhere (2 points) or from left (1 points)
		leftMin = min(leftMin+2*(int(s[i])-'0'), i+1)
		// find minimum in res and leftMin + right part [i+1, n)
		res = min(res, n-1-i+leftMin)
	}
	return res
}
