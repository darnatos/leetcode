package myLinkedList

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	count := 1
	p := head
	for p != nil {
		count++
		p = p.Next
	}

	if count-n == 1 {
		if head.Next != nil {
			head = head.Next
		} else {
			head = nil
		}
		return head
	}

	p = head
	q := head
	for m := count - n; m > 1; m-- {
		q = p
		if p.Next == nil {
			return head
		}
		p = p.Next
	}

	if q.Next == nil {
		q.Next = nil
	} else {
		q.Next = p.Next
	}
	q = nil

	return head
}
