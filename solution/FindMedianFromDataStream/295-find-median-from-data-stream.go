package FindMedianFromDataStream

import "container/heap"

type MinHeap []int

func (ih MinHeap) Len() int           { return len(ih) }
func (ih MinHeap) Less(i, j int) bool { return ih[i] < ih[j] }
func (ih MinHeap) Swap(i, j int)      { ih[i], ih[j] = ih[j], ih[i] }
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
	*ih = append(*ih, x.(int))
}

type MedianFinder struct {
	maxHeap *MinHeap
	minHeap *MinHeap
	even    bool
}

func Constructor() MedianFinder {
	mah := make(MinHeap, 0)
	mih := make(MinHeap, 0)
	mf := MedianFinder{maxHeap: &mah, minHeap: &mih, even: true}
	return mf
}

func (this *MedianFinder) AddNum(num int) {
	if this.even {
		heap.Push(this.maxHeap, -num)
		heap.Push(this.minHeap, -heap.Pop(this.maxHeap).(int))
	} else {
		heap.Push(this.minHeap, num)
		heap.Push(this.maxHeap, -heap.Pop(this.minHeap).(int))
	}
	this.even = !this.even
}

func (this *MedianFinder) FindMedian() float64 {
	if this.even {
		x, y := this.minHeap.Peep().(int), -this.maxHeap.Peep().(int)
		return float64(x+y) / 2
	} else {
		x := this.minHeap.Peep().(int)
		return float64(x)
	}
}
