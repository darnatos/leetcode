package myLinkedList

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	mA := make([]*ListNode, 0, 4)
	cur := headA
	for cur != nil {
		mA = append(mA, cur)
		cur = cur.Next
	}
	cur = headB
	mB := make([]*ListNode, 0, 4)
	for cur != nil {
		mB = append(mB, cur)
		cur = cur.Next
	}
	i, j := len(mA)-1, len(mB)-1
	for i >= 0 && j >= 0 {
		if mA[i] != mB[j] {
			break
		}
		i, j = i-1, j-1
	}
	if i < len(mA)-1 {
		return mA[i+1]
	}
	return nil
}
