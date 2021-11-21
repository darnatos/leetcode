package solution

import (
	"github.com/darnatos/leetcode/solution/myLinkedList"
	"github.com/darnatos/leetcode/solution/treeNode"
)

func sortedListToBST(head *myLinkedList.ListNode) *treeNode.TreeNode {
	return toBST(head, nil)
}

func toBST(head, tail *myLinkedList.ListNode) *treeNode.TreeNode {
	if head == tail {
		return nil
	}
	slow, fast := head, head
	for fast != tail && fast.Next != tail {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return &treeNode.TreeNode{
		Val:   slow.Val,
		Left:  toBST(head, slow),
		Right: toBST(slow.Next, tail),
	}
}
