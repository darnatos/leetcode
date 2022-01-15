package solution

func minJumps(arr []int) int {
	n := len(arr)
	m := make(map[int][]int)
	for i := range arr {
		m[arr[i]] = append(m[arr[i]], i)
	}

	dp := make([]int, n)
	visited := make([]bool, n)
	q := make([]int, 0, n)
	q = append(q, 0)
	visited[0] = true
	var top, step, size int
	for len(q) != 0 {
		size = len(q)
		for i := size - 1; i >= 0; i-- {
			top = q[i]
			if top == n-1 {
				return step
			}
			dp[top] = step
			for _, j := range m[arr[top]] {
				if visited[j] {
					continue
				}
				visited[j] = true
				q = append(q, j)
			}
			delete(m, arr[top])
			if top > 0 && !visited[top-1] {
				visited[top-1] = true
				q = append(q, top-1)
			}
			if top < n-1 && !visited[top+1] {
				visited[top+1] = true
				q = append(q, top+1)
			}
		}
		step++
		q = q[size:]
	}
	return 0
}
