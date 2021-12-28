package solution

func generateMatrix(n int) [][]int {
	nn := n * n
	dir := [5]int{0, 1, 0, -1, 0}
	d := 0
	x, y := 0, 0
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	for i := 1; i <= nn; i++ {
		res[x][y] = i
		if (d == 0 && (y == n-1 || res[x][y+1] != 0)) || d == 1 && (x == n-1 || res[x+1][y] != 0) || (d == 2 && (y == 0 || res[x][y-1] != 0)) || (d == 3 && (x == 0 || res[x-1][y] != 0)) {
			d++
		}
		if d >= 4 {
			d = 0
		}
		x, y = x+dir[d], y+dir[d+1]
	}
	return res
}
