package solution

import "sort"

func highestRankedKItems(grid [][]int, pricing []int, start []int, k int) [][]int {
	res := make([][]int, 0, k)

	dir := []int{0, -1, 0, 1, 0}
	q := make([][]int, 0, 1)
	q = append(q, start)
	step := 0
	for len(q) != 0 {
		v := len(q)
		for u := 0; u < v; u++ {
			cur := q[u]
			i, j := cur[0], cur[1]
			if grid[i][j] == 0 {
				continue
			}
			if grid[i][j] >= pricing[0] && grid[i][j] <= pricing[1] {
				res = append(res, []int{i, j, step + 1, grid[i][j]})
			}
			grid[i][j] = 0

			for d := 0; d < 4; d++ {
				ni, nj := i+dir[d], j+dir[d+1]
				if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[0]) && grid[ni][nj] > 0 {
					q = append(q, []int{ni, nj})
				}
			}
		}
		step++
		q = q[v:]
	}
	// fmt.Println(res)
	sort.Slice(res, func(i, j int) bool {
		if res[i][2] == res[j][2] {
			if res[i][3] == res[j][3] {
				if res[i][0] == res[j][0] {
					return res[i][1] < res[j][1]
				}
				return res[i][0] < res[j][0]
			}
			return res[i][3] < res[j][3]
		}
		return res[i][2] < res[j][2]
	})
	res = res[:min(len(res), k)]
	for i := range res {
		res[i] = res[i][:2]
	}
	return res
}
