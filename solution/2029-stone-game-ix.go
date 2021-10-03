package solution

func StoneGameIX(stones []int) bool {
	count := [3]int{}
	for _, v := range stones {
		count[v%3]++
	}

	less, greater := count[1], count[2]
	if less > greater {
		less, greater = greater, less
	}

	if less == 0 {
		return greater >= 3 && count[0]%2 > 0
	}
	return greater-less >= 3 || count[0]%2 == 0
}
