package myLinkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	cur := head
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				cur.Next = l1
				l1 = l1.Next
			} else {
				cur.Next = l2
				l2 = l2.Next
			}
			cur = cur.Next
		} else if l1 == nil {
			for l2 != nil {
				cur.Next = l2
				l2 = l2.Next
				cur = cur.Next
			}
		} else {
			for l1 != nil {
				cur.Next = l1
				l1 = l1.Next
				cur = cur.Next
			}
		}
	}
	return head.Next
}
