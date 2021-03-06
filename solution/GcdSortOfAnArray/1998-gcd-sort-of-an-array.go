package GcdSortOfAnArray

import (
	"github.com/darnatos/leetcode/solution/UnionFind"
	"math"
	"sort"
)

func GcdSort(nums []int) bool {
	maxNum := math.MinInt32
	for i := range nums {
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
	}
	spf := seize(maxNum)
	uf := UnionFind.Constructor(maxNum)

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
