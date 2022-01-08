package myLinkedList

import "math/rand"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	head  *ListNode
	count int
}

func LLRNConstructor(head *ListNode) Solution {
	count := 0
	for cur := head; cur != nil; cur = cur.Next {
		count++
	}
	return Solution{head, count}
}

func (this *Solution) GetRandom() int {
	cur := this.head
	for d := rand.Intn(this.count); d > 0; d-- {
		cur = cur.Next
	}
	return cur.Val
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
