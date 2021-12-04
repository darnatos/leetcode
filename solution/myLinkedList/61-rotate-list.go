package myLinkedList

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}
	cnt := 0
	cur := head
	for cur != nil {
		cnt++
		cur = cur.Next
	}
	target := ((cnt-k)%cnt + cnt) % cnt
	if target == 0 {
		return head
	}

	cur = head
	for i := 0; i < target-1; i++ {
		cur = cur.Next
	}

	tail := cur.Next
	for tail.Next != nil {
		tail = tail.Next
	}
	res := cur.Next
	cur.Next = nil
	tail.Next = head
	return res
}
