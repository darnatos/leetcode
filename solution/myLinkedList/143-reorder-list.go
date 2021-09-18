package myLinkedList

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	prev, slow, fast := head, head, head
	for fast != nil && fast.Next != nil {
		prev, slow, fast = slow, slow.Next, fast.Next.Next
	}
	prev.Next = nil

	rev := reverseList(slow)
	ptr := head

	for ptr != nil {
		t1, t2 := ptr.Next, rev.Next
		ptr.Next = rev
		if t1 != nil {
			rev.Next = t1
		}
		ptr, rev = t1, t2
	}
}
