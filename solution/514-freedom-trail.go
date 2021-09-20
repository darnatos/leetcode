package solution

import "math"

func FindRotateSteps(ring string, key string) int {
	m, n := len(ring), len(key)
	cMap := make(map[rune][]int, 26)

	for i, r := range ring {
		if _, ok := cMap[r]; !ok {
			cMap[r] = []int{i}
			continue
		}
		cMap[r] = append(cMap[r], i)
	}

	state := map[int]int{0: 0}

	for _, c := range key {
		nextState := make(map[int]int, len(cMap[c]))
		for _, j := range cMap[c] {
			nextState[j] = math.MaxInt32
			for i := range state {
				diff := abs(i - j)
				dist := min(diff, m-diff)
				nextState[j] = min(nextState[j], dist+state[i])
			}
		}
		state = nextState
	}

	res := math.MaxInt32
	for i := range state {
		res = min(res, state[i])
	}

	return res + n
}
