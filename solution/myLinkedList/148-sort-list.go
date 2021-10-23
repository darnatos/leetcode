package myLinkedList

func SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head
	length := 0
	for cur != nil {
		length++
		cur = cur.Next
	}

	sentinel := new(ListNode)
	sentinel.Next = head
	var left, right, tail *ListNode
	for step := 1; step < length; step = step << 1 {
		cur = sentinel.Next
		tail = sentinel
		for cur != nil {
			left = cur
			right = split(left, step)
			cur = split(right, step)
			tail = merge(left, right, tail)
		}
	}
	return sentinel.Next
}

func split(head *ListNode, step int) *ListNode {
	for i := 1; head != nil && i < step; i++ {
		head = head.Next
	}
	if head == nil {
		return nil
	}
	second := head.Next
	head.Next = nil
	return second
}

func merge(left, right, head *ListNode) *ListNode {
	cur := head
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next, cur, left = left, left, left.Next
		} else {
			cur.Next, cur, right = right, right, right.Next
		}
	}
	cur.Next = left
	if cur.Next == nil {
		cur.Next = right
	}
	for cur.Next != nil {
		cur = cur.Next
	}
	return cur
}
