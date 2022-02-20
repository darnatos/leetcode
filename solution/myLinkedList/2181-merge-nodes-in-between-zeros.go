package myLinkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
	dummy := &ListNode{}
	res := dummy
	sum := 0
	for cur := head; cur != nil; cur = cur.Next {
		sum += cur.Val
		if cur.Next != nil && cur.Next.Val == 0 {
			dummy.Next = &ListNode{Val: sum}
			dummy = dummy.Next
			cur, cur.Next = cur.Next, nil
			sum = 0
		}
	}
	return res.Next
}
