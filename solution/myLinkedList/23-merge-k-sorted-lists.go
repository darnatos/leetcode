package myLinkedList

import "container/heap"

type ListNodeHeap []*ListNode

func (h ListNodeHeap) Len() int { return len(h) }
func (h ListNodeHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}
func (h ListNodeHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *ListNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *ListNodeHeap) Pop() interface{} {
	p := *h
	el := p[len(p)-1]
	*h = p[0 : len(p)-1]
	return el
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &ListNodeHeap{}

	for i := range lists {
		if lists[i] != nil {
			heap.Push(h, lists[i])
		}
	}

	if h.Len() == 0 {
		return nil
	}

	first := heap.Pop(h).(*ListNode)
	head := first
	tail := first

	if first.Next != nil {
		heap.Push(h, first.Next)
	}

	for h.Len() != 0 {
		curr := heap.Pop(h).(*ListNode)

		tail.Next = curr
		tail = tail.Next

		if curr.Next != nil {
			heap.Push(h, curr.Next)
		}
	}

	return head
}
