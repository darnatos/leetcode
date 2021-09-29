package myLinkedList

func splitListToParts(head *ListNode, k int) []*ListNode {
	res := make([]*ListNode, k)
	count, cur := 0, head
	for i := 0; cur != nil; i++ {
		count++
		cur = cur.Next
	}

	l, r := count/k, count%k
	var pre *ListNode
	cur = head
	for i := 0; cur != nil && i < k; i++ {
		res[i] = cur
		for j := 0; j < l; j++ {
			pre, cur = cur, cur.Next
		}
		if r > 0 {
			pre, cur = cur, cur.Next
			r--
		}
		pre.Next = nil
	}

	return res
}
