package myLinkedList

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}


func NewListNode(vals ...int) *ListNode{
	head := new(ListNode)
	cur := head
	for i := range vals {
		cur.Val = vals[i]
		if i < len(vals)-1 {
			cur.Next = new(ListNode)
			cur = cur.Next
		}
	}
	return head
}

func (n *ListNode) Print(){
	for cur := n; cur != nil; cur = cur.Next {
		fmt.Printf("%v ", cur.Val)
	}
	fmt.Print("\r\n")
}