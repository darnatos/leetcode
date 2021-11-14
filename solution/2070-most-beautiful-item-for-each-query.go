package solution

import "sort"

func MaximumBeauty(items [][]int, queries []int) []int {
	sort.Slice(items, func(i, j int) bool {
		if items[i][0] == items[j][0] {
			return items[i][1] > items[j][1]
		}
		return items[i][0] < items[j][0]
	})
	m := make(map[int]int, len(items))
	for i := range items {
		if _, ok := m[items[i][0]]; !ok {
			m[items[i][0]] = items[i][1]
		}
	}
	keys := make([]int, 0, len(m))
	keys = append(keys, 0)
	for i := range m {
		keys = append(keys, i)
	}
	sort.Ints(keys)
	maxi := 0
	for _, k := range keys {
		if m[k] > maxi {
			maxi = m[k]
		} else {
			m[k] = maxi
		}
	}
	res := make([]int, len(queries))
	for i, q := range queries {
		l, r := 0, len(keys)-1
		for l <= r {
			mid := l + (r-l)/2
			if keys[mid] > q {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}

		res[i] = m[keys[l-1]]
	}

	return res
}
