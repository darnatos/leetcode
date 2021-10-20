package solution

import (
	"container/heap"
	"sort"
)

type Skyline struct {
	l, r, h int
	idx     int
}

type SLHeap []*Skyline

func (hp SLHeap) Len() int           { return len(hp) }
func (hp SLHeap) Less(i, j int) bool { return hp[i].h > hp[j].h }
func (hp SLHeap) Swap(i, j int) {
	hp[i], hp[j] = hp[j], hp[i]
	hp[i].idx, hp[j].idx = hp[j].idx, hp[i].idx
}
func (hp *SLHeap) Push(x interface{}) {
	n := hp.Len()
	skyline := x.(*Skyline)
	skyline.idx = n
	*hp = append(*hp, skyline)
}
func (hp *SLHeap) Pop() interface{} {
	n := hp.Len()
	old := *hp
	x := old[n-1]
	*hp = old[:n-1]

	return x
}

func GetSkyline(buildings [][]int) [][]int {
	m := make(map[int][]*Skyline, len(buildings)<<1)
	for _, b := range buildings {
		skyline := &Skyline{
			l: b[0],
			r: b[1],
			h: b[2],
		}

		m[b[0]] = append(m[b[0]], skyline)
		m[b[1]] = append(m[b[1]], skyline)
	}

	positions := make([]int, 0, len(m))
	for pos := range m {
		positions = append(positions, pos)
	}
	sort.Ints(positions)

	hp := SLHeap{&Skyline{l: -1, r: -1, h: 0}}
	silhouette := make([][]int, 0, len(positions))

	for _, pos := range positions {
		curHeight := hp[0].h

		for _, skyline := range m[pos] {
			if skyline.r > pos {
				heap.Push(&hp, skyline)
			} else {
				heap.Remove(&hp, skyline.idx)
			}
		}

		if hp[0].h != curHeight {
			silhouette = append(silhouette, []int{pos, hp[0].h})
		}
	}

	return silhouette
}
