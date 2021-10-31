package solution

func MinimumOperations(nums []int, start int, goal int) int {
	step := 0
	visited := make(map[int]bool, 1001)
	queue := []int{start}
	for len(queue) != 0 {
		n := len(queue)
		step++
		for i := 0; i < n; i++ {
			x := queue[i]
			for _, v := range nums {
				tmp := x + v
				if tmp == goal {
					return step
				}
				if !visited[tmp] && tmp >= 0 && tmp <= 1000 {
					visited[tmp] = true
					queue = append(queue, tmp)
				}
				tmp = x - v
				if tmp == goal {
					return step
				}
				if !visited[tmp] && tmp >= 0 && tmp <= 1000 {
					visited[tmp] = true
					queue = append(queue, tmp)
				}
				tmp = x ^ v
				if tmp == goal {
					return step
				}
				if !visited[tmp] && tmp >= 0 && tmp <= 1000 {
					visited[tmp] = true
					queue = append(queue, tmp)
				}
			}
		}
		queue = queue[n:]
	}
	return -1
}
