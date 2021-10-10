package StockPrice

import (
	"container/heap"
)

type TimePrice struct {
	time  int
	price int
}
type MinHeap []TimePrice

func (ih MinHeap) Len() int           { return len(ih) }
func (ih MinHeap) Less(i, j int) bool { return ih[i].price < ih[j].price }
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
	*ih = append(*ih, x.(TimePrice))
}

type StockPrice struct {
	latestTime int
	timePrices map[int]int
	minPrices  MinHeap
	maxPrices  MinHeap
}

func Constructor() StockPrice {
	minPrices := make(MinHeap, 0)
	maxPrices := make(MinHeap, 0)
	return StockPrice{0, make(map[int]int), minPrices, maxPrices}
}

func (this *StockPrice) Update(timestamp int, price int) {
	this.timePrices[timestamp] = price
	this.latestTime = max(this.latestTime, timestamp)
	heap.Push(&this.minPrices, TimePrice{timestamp, price})
	heap.Push(&this.maxPrices, TimePrice{timestamp, -price})
}

func (this *StockPrice) Current() int {
	return this.timePrices[this.latestTime]
}

func (this *StockPrice) Maximum() int {
	timePrice := heap.Pop(&this.maxPrices).(TimePrice)
	for this.timePrices[timePrice.time] != -timePrice.price {
		timePrice = heap.Pop(&this.maxPrices).(TimePrice)
	}
	heap.Push(&this.maxPrices, timePrice)
	return -timePrice.price
}

func (this *StockPrice) Minimum() int {
	timePrice := heap.Pop(&this.minPrices).(TimePrice)
	for this.timePrices[timePrice.time] != timePrice.price {
		timePrice = heap.Pop(&this.minPrices).(TimePrice)
	}
	heap.Push(&this.minPrices, timePrice)
	return timePrice.price
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */
