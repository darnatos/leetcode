package SORTracker

import "container/heap"

type Pair struct {
	name  string
	score int
}

type MaxHeap []Pair

func (ih MaxHeap) Len() int { return len(ih) }
func (ih MaxHeap) Less(i, j int) bool {
	return ih[i].score > ih[j].score || ih[i].score == ih[j].score && ih[i].name < ih[j].name
}
func (ih MaxHeap) Swap(i, j int) { ih[i], ih[j] = ih[j], ih[i] }
func (ih *MaxHeap) Pop() interface{} {
	old := *ih
	n := len(old)
	x := old[n-1]
	*ih = old[:n-1]
	return x
}
func (ih *MaxHeap) Peep() interface{} {
	o := *ih
	return o[0]
}

func (ih *MaxHeap) Push(x interface{}) {
	*ih = append(*ih, x.(Pair))
}

type MinHeap []Pair

func (ih MinHeap) Len() int { return len(ih) }
func (ih MinHeap) Less(i, j int) bool {
	return ih[i].score < ih[j].score || ih[i].score == ih[j].score && ih[i].name > ih[j].name
}
func (ih MinHeap) Swap(i, j int) { ih[i], ih[j] = ih[j], ih[i] }
func (ih *MinHeap) Pop() interface{} {
	old := *ih
	n := len(old)
	x := old[n-1]
	*ih = old[:n-1]
	return x
}
func (ih *MinHeap) Peep() interface{} {
	o := *ih
	return o[0]
}
func (ih *MinHeap) Push(x interface{}) {
	*ih = append(*ih, x.(Pair))
}

type SORTracker struct {
	pq1 *MaxHeap
	pq2 *MinHeap
}

func Constructor() SORTracker {
	pq1 := make(MaxHeap, 0)
	pq2 := make(MinHeap, 0)
	return SORTracker{&pq1, &pq2}
}

func (st *SORTracker) Add(name string, score int) {
	p := Pair{name, score}
	if st.pq2.Len() != 0 {
		top := st.pq2.Peep().(Pair)
		if p.score > top.score || p.score == top.score && p.name < top.name {
			heap.Push(st.pq1, heap.Pop(st.pq2))
			heap.Push(st.pq2, p)
			return
		}
	}
	heap.Push(st.pq1, p)
}

func (st *SORTracker) Get() string {
	tmp := heap.Pop(st.pq1)
	heap.Push(st.pq2, tmp)
	return tmp.(Pair).name
}
