package solution

/**
 * Definition for a Node.
 */

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}
	root := &Node{
		Val: node.Val,
	}
	m := make(map[int]*Node)
	m[node.Val] = root

	q := make([]*Node, 0)
	q = append(q, node)

	for len(q) > 0 {
		end := len(q)
		for i := 0; i < end; i++ {
			u := q[i]
			for _, v := range u.Neighbors {
				if _, ok := m[v.Val]; !ok {
					nn := &Node{
						Val: v.Val,
					}
					m[v.Val] = nn
					q = append(q, v)
				}
				m[u.Val].Neighbors = append(m[u.Val].Neighbors, m[v.Val])
			}
		}
		q = q[end:]
	}

	return root
}
