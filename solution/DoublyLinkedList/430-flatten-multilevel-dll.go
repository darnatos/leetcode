package DoublyLinkedList

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	cur := root

	for cur != nil {
		if cur.Child == nil {
			cur = cur.Next
			continue
		}

		child := cur.Child
		for child.Next != nil {
			child = child.Next
		}
		if cur.Next != nil {
			child.Next, cur.Next.Prev = cur.Next, child
		}
		cur.Next, cur.Child, cur.Child.Prev = cur.Child, nil, cur
		cur = cur.Next
	}

	return root
}
