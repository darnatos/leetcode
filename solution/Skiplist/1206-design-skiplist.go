package Skiplist

import (
	"math/rand"
	"time"
)

type Node struct {
	Val        int
	Next, Down *Node
}

type Skiplist struct {
	head *Node
	r    *rand.Rand
}

func Constructor() Skiplist {
	head := &Node{
		Val: -1,
	}
	return Skiplist{
		head: head,
		r:    rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (sl Skiplist) Search(target int) bool {
	cur := sl.head
	for cur != nil {
		for cur.Next != nil && cur.Next.Val < target {
			cur = cur.Next
		}
		if cur.Next != nil && cur.Next.Val == target {
			return true
		}
		cur = cur.Down
	}
	return false
}

func (sl *Skiplist) Add(num int) {
	stack := make([]*Node, 0, 16)
	cur := sl.head
	for cur != nil {
		for cur.Next != nil && cur.Next.Val < num {
			cur = cur.Next
		}
		stack = append(stack, cur)
		cur = cur.Down
	}
	insert := true
	var down *Node
	for insert && len(stack) != 0 {
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur.Next = &Node{
			Val:  num,
			Next: cur.Next,
			Down: down,
		}
		down = cur.Next
		insert = sl.r.Intn(5000) < 2500
	}
	if insert {
		sl.head = &Node{
			Val:  -1,
			Down: sl.head,
		}
	}
}

func (sl *Skiplist) Erase(num int) bool {
	cur := sl.head
	found := false
	for cur != nil {
		for cur.Next != nil && cur.Next.Val < num {
			cur = cur.Next
		}
		if cur.Next != nil && cur.Next.Val == num {
			found = true
			cur.Next = cur.Next.Next
		}
		cur = cur.Down
	}
	return found
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
