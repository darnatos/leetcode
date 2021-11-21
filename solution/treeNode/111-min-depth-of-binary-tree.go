package treeNode

import (
	"container/list"
)

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	q := list.New()
	q.PushBack(root)
	level := 0
	for q.Len() != 0 {
		level++
		k := q.Len()
		for i := 0; i < k; i++ {
			node := q.Front().Value.(*TreeNode)
			q.Remove(q.Front())
			if node.Right == nil && node.Left == nil {
				return level
			}
			if node.Right != nil {
				q.PushBack(node.Right)
			}
			if node.Left != nil {
				q.PushBack(node.Left)
			}
		}
	}

	return -1
}
