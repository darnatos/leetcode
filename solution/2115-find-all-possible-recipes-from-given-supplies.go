package solution

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	adj := make(map[string][]string)
	in := make(map[string]int, len(recipes))
	for i, r := range recipes {
		in[r] = len(ingredients[i])
		for _, j := range ingredients[i] {
			adj[j] = append(adj[j], r)
		}
	}
	q := supplies
	res := make([]string, 0, len(recipes))

	for len(q) != 0 {
		u := q[0]
		q = q[1:]
		for _, v := range adj[u] {
			in[v]--
			if in[v] == 0 {
				res = append(res, v)
				q = append(q, v)
			}
		}
	}

	return res
}
