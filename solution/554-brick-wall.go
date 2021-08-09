package solution

func LeastBrick(wall [][]int) int {
	m := make(map[int]int)
	edge := 0

	for _, row := range wall {
		b := 0
		for _, x := range row {
			b += x
			if _, ok := m[b]; ok {
				m[b]++
			} else {
				m[b] = 1
			}
		}

		if edge < b {
			edge = b
		}
	}

	max := 0
	for i, v := range m {
		if max < v && i != edge {
			max = v
		}
	}

	return len(wall) - max
}
