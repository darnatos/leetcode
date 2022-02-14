package solution

func groupStrings(words []string) []int {
	n := len(words)
	vertex := make(map[int]int, n)
	masks := make([]int, 0, n)
	var mask int
	for _, w := range words {
		mask = 0
		for j := range w {
			mask |= 1 << (w[j] - 'a')
		}
		vertex[mask]++
		masks = append(masks, mask)
	}

	var cnt int
	res := []int{0, 0}
	for u := range masks {
		if _, ok := vertex[masks[u]]; !ok {
			continue
		}
		res[0]++
		cnt = 0
		stack := make([]int, 0)
		stack = append(stack, masks[u])
		for len(stack) > 0 {
			mask = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			cnt += vertex[mask]
			delete(vertex, mask)
			for i := 0; i < 26; i++ {
				tmp := mask ^ (1 << i)
				if _, ok := vertex[tmp]; ok {
					stack = append(stack, tmp)
				}
				for j := i + 1; j < 26; j++ {
					if mask>>i&1 != mask>>j&1 {
						if _, ok := vertex[tmp^(1<<j)]; ok {
							stack = append(stack, tmp^(1<<j))
						}
					}
				}
			}
		}
		if res[1] < cnt {
			res[1] = cnt
		}
	}
	return res
}
