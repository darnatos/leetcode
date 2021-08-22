package myLinkedList

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	res0 := &ListNode{0, nil}
	res := res0
	carry := false

	for l2 != nil || l1 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if carry {
			sum += 1
			carry = false
		}
		if sum >= 10 {
			sum -= 10
			carry = true
		}
		res.Val = sum
		if l2 != nil || l1 != nil {
			res.Next = &ListNode{0, nil}
		} else {
			if carry {
				res.Next = &ListNode{1, nil}
			}
		}
		res = res.Next
	}

	return res0
}
