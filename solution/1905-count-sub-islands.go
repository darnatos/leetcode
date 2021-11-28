package solution

func CountSubIslands(grid1 [][]int, grid2 [][]int) int {
	res := 0
	for i := range grid2 {
		for j := range grid2[0] {
			if grid2[i][j] == 0 {
				continue
			}
			if csiDfs(grid1, grid2, i, j) {
				res++
			}
		}
	}
	return res
}

func csiDfs(grid1, grid2 [][]int, i, j int) bool {
	if i < 0 || i > len(grid2)-1 || j < 0 || j > len(grid2[0])-1 || grid2[i][j] == 0 {
		return true
	}
	grid2[i][j] = 0
	res := csiDfs(grid1, grid2, i+1, j)
	res = csiDfs(grid1, grid2, i-1, j) && res
	res = csiDfs(grid1, grid2, i, j+1) && res
	res = csiDfs(grid1, grid2, i, j-1) && res
	return grid1[i][j] == 1 && res
}
