package solution

import "sort"

func FindAllPeople(n int, meetings [][]int, firstPerson int) []int {
	sort.Slice(meetings, func(i, j int) bool { return meetings[i][2] < meetings[j][2] })
	ds := make([]int, n)
	for i := range ds {
		ds[i] = i
	}
	union(ds, 0, firstPerson)
	for i := 0; i < len(meetings); {
		curTime := meetings[i][2]
		pool := make([]int, 0, 8)
		for ; i < len(meetings) && meetings[i][2] == curTime; i++ {
			union(ds, meetings[i][0], meetings[i][1])
			pool = append(pool, meetings[i][0])
			pool = append(pool, meetings[i][1])
		}
		for _, i := range pool {
			if find(ds, i) != find(ds, 0) {
				ds[i] = i
			}
		}
	}
	res := make([]int, 0, 4)
	for i := 0; i < n; i++ {
		if find(ds, i) == find(ds, 0) {
			res = append(res, i)
		}
	}
	return res
}
