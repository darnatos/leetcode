package myLinkedList

func ReverseKGroup(head *ListNode, k int) *ListNode {
	cur := head
	var tail *ListNode
	for i := 0; i < k; i++ {
		if cur == nil {
			return head
		}
		cur = cur.Next
	}
	head, tail = reverseK(head, k)
	cnt := 0
	for cur != nil {
		cnt++
		cur = cur.Next
		if cnt == k {
			tail.Next, tail = reverseK(tail.Next, k)
			cnt = 0
		}
	}

	return head
}

func reverseK(h *ListNode, k int) (*ListNode, *ListNode) {
	cur := h
	var prev *ListNode

	for k > 0 && cur != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
		k--
	}
	h.Next = cur
	return prev, h
}
