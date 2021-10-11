package solution

import (
	"math"
	"strings"
)

func ShortestSuperstring(words []string) string {
	n := len(words)
	arr := make([]int, n*n)
	graph := make([][]int, n)
	for i := range graph {
		graph[i] = arr[i*n : (i+1)*n]
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			graph[i][j] = calcStrDist(words[i], words[j])
			graph[j][i] = calcStrDist(words[j], words[i])
		}
	}
	arrDp := make([]int, (1<<n)*n)
	arrPath := make([]int, (1<<n)*n)
	dp := make([][]int, 1<<n)
	path := make([][]int, 1<<n)
	min, last := math.MaxInt32, -1

	for i := 1; i < 1<<n; i++ {
		dp[i] = arrDp[i*n : (i+1)*n]
		path[i] = arrPath[i*n : (i+1)*n]
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt32
		}
		for j := 0; j < n; j++ {
			if i&(1<<j) > 0 {
				prev := i - (1 << j)
				if prev == 0 {
					dp[i][j] = len(words[j])
				} else {
					for k := 0; k < n; k++ {
						if dp[prev][k] < math.MaxInt32 && dp[prev][k]+graph[k][j] < dp[i][j] {
							dp[i][j] = dp[prev][k] + graph[k][j]
							path[i][j] = k
						}
					}
				}
			}
			if i == 1<<n-1 && dp[i][j] < min {
				min = dp[i][j]
				last = j
			}
		}
	}
	sb := make([]byte, 0)
	cur := 1<<n - 1
	stack := make([]int, 0, n)
	for cur > 0 {
		stack = append(stack, last)
		tmp := cur
		cur -= 1 << last
		last = path[tmp][last]
	}
	//fmt.Println(graph, dp, path, stack)
	i := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	sb = append(sb, []byte(words[i])...)
	for len(stack) > 0 {
		j := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		sb = append(sb, []byte(words[j][len(words[j])-graph[i][j]:])...)
		i = j
	}

	return string(sb)
}

func calcStrDist(a, b string) int {
	for i := 1; i < len(a); i++ {
		if strings.Index(b, a[i:]) == 0 {
			return len(b) - len(a) + i
		}
	}
	return len(b)
}
