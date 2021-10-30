package solution

import (
	"sort"
)

func MaxTwoEvents(events [][]int) int {
	n := len(events)

	es := make([][]int, 2*n)
	for i := range events {
		es[i] = []int{events[i][0], events[i][2], 1}
		es[i+n] = []int{events[i][1], events[i][2], -1}
	}
	//fmt.Println(es)
	sort.Slice(es, func(i, j int) bool {
		if es[i][0] != es[j][0] {
			return es[i][0] < es[j][0]
		}
		return es[i][2] > es[j][2]
	})
	//fmt.Println(es)
	mx := 0
	res := 0
	for _, e := range es {
		//fmt.Println(e, mx, res)
		if e[2] == -1 {
			mx = max(mx, e[1])
		} else {
			res = max(res, mx+e[1])
		}
	}

	return res
}
