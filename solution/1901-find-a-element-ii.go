package solution

func FindPeakGrid(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	lo, hi := 0, n-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		maxRow := 0
		for i := 0; i < m; i++ {
			if mat[maxRow][mid] <= mat[i][mid] {
				maxRow = i
			}
		}
		leftIsBig := mid > lo && mat[maxRow][mid-1] > mat[maxRow][mid]
		rightIsBig := hi > mid && mat[maxRow][mid+1] > mat[maxRow][mid]

		if !leftIsBig && !rightIsBig {
			return []int{maxRow, mid}
		} else if rightIsBig {
			lo = mid + 1
		} else {
			hi = mid - 1
		}

	}
	return []int{-1, -1}
}
