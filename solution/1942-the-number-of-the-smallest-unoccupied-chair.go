package solution

import (
	"container/heap"
	"sort"
)

// QItem represents the interface to be implemented stored in this queue
type QItem interface {
	Less(item QItem) bool
}

// priorityQueueImpl for the underlying implementation of priority queues
type priorityQueueImpl []QItem

// Len get queue length
func (pqi priorityQueueImpl) Len() int {
	return len(pqi)
}

// Less is used for element comparison
func (pqi priorityQueueImpl) Less(i, j int) bool {
	return pqi[i].Less(pqi[j])
}

func (pqi priorityQueueImpl) Swap(i, j int) {
	pqi[i], pqi[j] = pqi[j], pqi[i]
}

// Push is used to push an object into the queue
func (pqi *priorityQueueImpl) Push(x interface{}) {
	item := x.(QItem)
	*pqi = append(*pqi, item)
}

// Pop pops an object out of the queue
func (pqi *priorityQueueImpl) Pop() interface{} {
	old := *pqi
	n := len(old)
	item := old[n-1]
	*pqi = old[0 : n-1]
	return item
}

// PriorityQueue implements priority queue
type PriorityQueue struct {
	priorityQueueImpl
}

// NewPriorityQueue is used to build PriorityQueue
func NewPriorityQueue() *PriorityQueue {
	var pq PriorityQueue
	heap.Init(&pq.priorityQueueImpl)
	return &pq
}

// Push is used to push an object into the queue
func (pq *PriorityQueue) Push(item QItem) {
	heap.Push(&pq.priorityQueueImpl, item)
}

// Pop is used to pop an object from the queue
func (pq *PriorityQueue) Pop() QItem {
	return heap.Pop(&pq.priorityQueueImpl).(QItem)
}

// Front is used to get the minimum value in the current queue
func (pq *PriorityQueue) Front() QItem {
	// The first bit in the queue should be the minimum
	return pq.priorityQueueImpl[0]
}

// Length is used to get the length of the current queue
func (pq *PriorityQueue) Length() int {
	return pq.priorityQueueImpl.Len()
}

type Item struct {
	num   int
	value *int
	time  []int
}

func (a Item) Less(b QItem) bool {
	return *(a.value) < *(b.(Item).value)
}

type Times [][]int

func (t Times) Len() int {
	return len(t)
}
func (t Times) Less(i, j int) bool {
	return t[i][0] < t[j][0]
}
func (t Times) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func SmallestChair(times [][]int, targetFriend int) int {
	targetFriendArrival := times[targetFriend][0]
	sort.Sort(Times(times))

	inUsePq := NewPriorityQueue()
	emptyPq := NewPriorityQueue()
	firstNewEmpty := 0

	for _, time := range times {
		for inUsePq.Len() > 0 && inUsePq.Front().(Item).time[1] <= time[0] {
			item := inUsePq.Pop().(Item)
			item.value = &item.num
			emptyPq.Push(item)
		}
		var seat Item
		if emptyPq.Length() == 0 {
			seat = Item{firstNewEmpty, &time[1], time}
			firstNewEmpty++
		} else {
			seat = emptyPq.Pop().(Item)
			seat.time = time
		}
		if seat.time[0] == targetFriendArrival {
			return seat.num
		}
		seat.value = &seat.time[1]
		inUsePq.Push(seat)
	}

	return -1
}
