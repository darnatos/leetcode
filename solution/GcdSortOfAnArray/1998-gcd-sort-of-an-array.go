package GcdSortOfAnArray

import (
	"math"
	"sort"
)

type UnionFind struct {
	parent []int
}

func Constructor(n int) UnionFind {
	parent := make([]int, n+1)
	for i := range parent {
		parent[i] = i
	}
	uf := UnionFind{parent: parent}
	return uf
}

func (uf UnionFind) Find(u int) int {
	if uf.parent[u] == u {
		return u
	}
	uf.parent[u] = uf.Find(uf.parent[u])
	return uf.parent[u]
}

func (uf UnionFind) Union(u, v int) {
	p, q := uf.Find(u), uf.Find(v)
	if p != q {
		uf.parent[p] = q
	}
}

func GcdSort(nums []int) bool {
	maxNum := math.MinInt32
	for i := range nums {
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
	}
	spf := seize(maxNum)
	uf := Constructor(maxNum)

	for _, v := range nums {
		for f := v; f > 1; f /= spf[f] {
			uf.Union(v, spf[f])
		}
	}

	sortedNums := append([]int{}, nums...)
	sort.Ints(sortedNums)

	for i := range nums {
		if uf.Find(nums[i]) != uf.Find(sortedNums[i]) {
			return false
		}
	}
	return true
}

func seize(n int) []int {
	spf := make([]int, n+1)
	for i := 2; i <= n; i++ {
		spf[i] = i
	}
	for i := 2; i*i <= n; i++ {
		if spf[i] != i {
			continue
		}
		for j := i * i; j <= n; j += i {
			if spf[j] > i {
				spf[j] = i
			}
		}
	}
	return spf
}
