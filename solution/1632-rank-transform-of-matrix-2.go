package solution

import (
	"math"
	"sort"
)

func unionFind(ds []int, i int) int {
	if ds[i] < 0 {
		return i
	}
	ds[i] = unionFind(ds, ds[i])
	return ds[i]
}

func updateRank(m [][]int, crds [][2]int, cols, rows []int) {
	ds := make([]int, len(m)+len(m[0]))
	for i := range ds {
		ds[i] = -1
	}

	for _, crd := range crds {
		pi, pj := unionFind(ds, crd[0]), unionFind(ds, len(rows)+crd[1])
		if pi != pj {
			ds[pi] = min(ds[pi], min(ds[pj], -max(rows[crd[0]], cols[crd[1]])-1))
			ds[pj] = pi
		}
	}

	for _, crd := range crds {
		cols[crd[1]] = -ds[unionFind(ds, crd[0])]
		rows[crd[0]] = cols[crd[1]]
		m[crd[0]][crd[1]] = rows[crd[0]]
	}
}

func MatrixRankTransform2(m [][]int) [][]int {
	si, sj, lastVal := len(m), len(m[0]), math.MinInt32
	mm := make([][3]int, 0, si*sj)
	rows := make([]int, si)
	cols := make([]int, sj)
	for i := 0; i < si; i++ {
		for j := 0; j < sj; j++ {
			mm = append(mm, [3]int{m[i][j], i, j})
		}
	}
	sort.Slice(mm, func(a, b int) bool {
		return mm[a][0] < mm[b][0]
	})

	crds := make([][2]int, 0, len(mm))
	for _, v := range mm {
		if v[0] != lastVal {
			updateRank(m, crds, cols, rows)
			lastVal = v[0]
			crds = crds[:0]
		}
		crds = append(crds, [2]int{v[1], v[2]})
	}

	updateRank(m, crds, cols, rows)
	return m
}
