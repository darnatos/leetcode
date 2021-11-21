package RangeFrequencyQueries

import "sort"

type RangeFreqQuery struct {
	m map[int][]int
}

func Constructor(arr []int) RangeFreqQuery {
	m := make(map[int][]int, len(arr))
	for i := range arr {
		m[arr[i]] = append(m[arr[i]], i)
	}
	for k := range m {
		sort.Ints(m[k])
	}
	return RangeFreqQuery{
		m: m,
	}
}

func (rfq *RangeFreqQuery) Query(left int, right int, value int) int {
	arr := rfq.m[value]
	return lowerBound(arr, right+1) - lowerBound(arr, left)
}

func lowerBound(arr []int, target int) int {
	lo, hi := 0, len(arr)-1
	for hi >= lo {
		mid := (lo + hi) / 2
		if arr[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return lo
}

/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */
