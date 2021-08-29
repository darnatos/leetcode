package solution

import (
	"sort"
	"sync"
)

func FindItinerary(tickets [][]string) []string {
	l := len(tickets) + 1
	adj := make(map[string][]string)
	for _, v := range tickets {
		adj[v[0]] = append(adj[v[0]], v[1])
	}

	wg := sync.WaitGroup{}
	for k := range adj {
		wg.Add(1)
		go func(k string) {
			sort.Strings(adj[k])
			wg.Done()
		}(k)
	}
	wg.Wait()

	res := make([]string, 0, l)

	dfs(adj, "JFK", &res)
	i, j := 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}

func dfs(adj map[string][]string, from string, res *[]string) {
	for len(adj[from]) > 0 {
		to := adj[from][0]
		adj[from] = adj[from][1:]
		dfs(adj, to, res)
	}
	*res = append(*res, from)
}
