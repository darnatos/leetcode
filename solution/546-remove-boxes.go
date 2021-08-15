package solution

var dpTable [][][]int

func RemoveBoxes(boxes []int) int {
	dpTable = *new3DArray(len(boxes))
	return dp(boxes, 0, len(boxes)-1, 0)
}

func dp(boxes []int, l, r, k int) int {
	if l > r {
		return 0
	}
	if dpTable[l][r][k] > 0 {
		return dpTable[l][r][k]
	}
	l0, k0 := l, k

	for l+1 <= r && boxes[l] == boxes[l+1] {
		l++
		k++
	}
	points := (k+1)*(k+1) + dp(boxes, l+1, r, 0)

	for m := l + 1; m <= r; m++ {
		if boxes[l] == boxes[m] {
			points = max(points, dp(boxes, m, r, k+1)+dp(boxes, l+1, m-1, 0))
		}
	}

	dpTable[l0][r][k0] = points
	return points
}

func new3DArray(n int) *[][][]int {
	t := make([][][]int, n)
	for i := 0; i < n; i++ {
		t[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			t[i][j] = make([]int, n)
		}
	}
	return &t
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
