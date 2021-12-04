package solution

import (
	"sort"
)

type DisjointSet map[int]int

func (ds DisjointSet) Find(u int) int {
	if u == ds[u] {
		return u
	}
	ds[u] = ds.Find(ds[u])
	return ds[u]
}

func (ds DisjointSet) Union(u, v int) {
	if ds[u] == 0 {
		ds[u] = u
	}
	if ds[v] == 0 {
		ds[v] = v
	}
	pu, pv := ds.Find(u), ds.Find(v)
	if pu != pv {
		ds[pu] = pv
	}
}

func (ds DisjointSet) GetGroups() map[int][]int {
	groups := make(map[int][]int)
	for u := range ds {
		fu := ds.Find(u)
		groups[fu] = append(groups[fu], u)
	}
	return groups
}

func MatrixRankTransform(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	groupByValue := make(map[int][][2]int)
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			v := matrix[r][c]
			groupByValue[v] = append(groupByValue[v], [2]int{r, c})
		}
	}
	sortedValues := make([]int, len(groupByValue))
	i := 0
	for v := range groupByValue {
		sortedValues[i] = v
		i++
	}
	sort.Ints(sortedValues)

	rank := make([]int, m+n)
	for _, v := range sortedValues {
		cells := groupByValue[v]
		uf := make(DisjointSet)
		for _, cell := range cells {
			uf.Union(cell[0], cell[1]+m)
		}

		for _, group := range uf.GetGroups() {
			maxRank := 0
			for _, i := range group {
				maxRank = max(maxRank, rank[i])
			}
			maxRank += 1
			for _, i := range group {
				rank[i] = maxRank
			}
		}
		for _, cell := range cells {
			matrix[cell[0]][cell[1]] = rank[cell[0]]
		}
	}
	return matrix
}
