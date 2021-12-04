package solution

import (
	"github.com/darnatos/leetcode/solution/UnionFind"
)

func FindRedundantConnection(edges [][]int) []int {
	n := len(edges)
	uf := UnionFind.Constructor(n)

	for _, e := range edges {
		pu, pv := uf.Find(e[0]), uf.Find(e[1])
		if pu == pv {
			return e
		}
		uf.Union(e[0], e[1])
	}

	return nil
}
