package solution

import (
	"math"
	"strconv"
)

func CrackSafe(n, k int) string {
	seq := ""
	goal := int(math.Pow(float64(k), float64(n)))
	visited := make(map[string]bool, goal)
	for i := 0; i < n; i++ {
		seq += "0"
	}
	visited[seq] = true

	crackSafeDfs(&seq, goal, n, k, visited)
	return seq
}

func crackSafeDfs(seq *string, goal, n, k int, visited map[string]bool) bool {
	if len(*seq) == goal {
		return true
	}

	prev := (*seq)[len(*seq)-n+1:]
	for i := 0; i < k; i++ {
		nc := strconv.Itoa(i)
		next := prev + nc
		if !visited[next] {
			visited[next] = true
			*seq += nc
			if crackSafeDfs(seq, goal, n, k, visited) {
				return true
			} else {
				delete(visited, next)
				*seq = (*seq)[:len(*seq)-1]
			}
		}
	}
	return false
}
