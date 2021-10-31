package myLinkedList

import "math"

func nodesBetweenCriticalPoints(head *ListNode) []int {
	pos := make([]int, 0)
	prev := head
	cur := head.Next
	curInd := 1
	for cur.Next != nil {
		if prev.Val > cur.Val && cur.Next.Val > cur.Val || prev.Val < cur.Val && cur.Next.Val < cur.Val {
			pos = append(pos, curInd)
		}
		prev, cur, curInd = cur, cur.Next, curInd+1
	}
	res := []int{-1, -1}
	if len(pos) < 2 {
		return res
	}
	res[0] = math.MaxInt32
	for i := 1; i < len(pos); i++ {
		if res[0] > pos[i]-pos[i-1] {
			res[0] = pos[i] - pos[i-1]
		}
	}
	res[1] = pos[len(pos)-1] - pos[0]
	return res
}
