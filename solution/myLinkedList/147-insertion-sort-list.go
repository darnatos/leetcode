package myLinkedList

func insertionSortList(head *ListNode) *ListNode {
	res := new(ListNode)
	cur := head

	for cur != nil {
		prev := res
		next := res.Next
		for next != nil {
			if cur.Val < next.Val {
				break
			}
			prev = prev.Next
			next = next.Next
		}
		tmp := cur.Next
		cur.Next = next
		prev.Next = cur

		cur = tmp
	}

	return res.Next
}
