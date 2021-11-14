package myLinkedList

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	idx := 0
	gc := 1
	prev := new(ListNode)
	prev.Next = head
	for cur := head; cur != nil; cur = cur.Next {
		idx++
		if idx == gc || cur.Next == nil {
			if idx%2 == 0 {
				tmp := cur.Next
				cur = prev.Next
				prev.Next = reverseKE(prev.Next, gc)
				cur.Next = tmp
			}
			prev = cur
			idx = 0
			gc++
		}
	}

	return head
}

func reverseKE(head *ListNode, k int) *ListNode {
	var prev *ListNode

	for k > 0 && head != nil {
		head.Next, prev, head = prev, head, head.Next
		k--
	}

	return prev
}
