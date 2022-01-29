package myLinkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbersII(l1 *ListNode, l2 *ListNode) *ListNode {
	m, n := 0, 0
	for cur := l1; cur != nil; cur = cur.Next {
		m++
	}
	for cur := l2; cur != nil; cur = cur.Next {
		n++
	}
	if m < n {
		m, n, l1, l2 = n, m, l2, l1
	}
	stack := make([]*ListNode, 0)
	cur1 := l1
	for m > n {
		stack = append(stack, cur1)
		cur1 = cur1.Next
		m--
	}

	hasCarry := false
	for cur2 := l2; cur1 != nil; cur1, cur2 = cur1.Next, cur2.Next {
		cur1.Val += cur2.Val
		stack = append(stack, cur1)
		if cur1.Val > 9 {
			hasCarry = true
		}
	}
	carry := 0
	if hasCarry {
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			sum := (top.Val + carry) % 10
			carry = (top.Val + carry) / 10
			top.Val = sum
			stack = stack[:len(stack)-1]
		}
	}
	if carry > 0 {
		return &ListNode{Val: 1, Next: l1}
	}
	return l1
}
