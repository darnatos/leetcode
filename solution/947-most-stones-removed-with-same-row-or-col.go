package solution

func RemoveStones(stones [][]int) int {
	sr, sc := make(map[int][]int), make(map[int][]int)

	for i := range stones {
		sr[stones[i][0]] = append(sr[stones[i][0]], stones[i][1])
		sc[stones[i][1]] = append(sc[stones[i][1]], stones[i][0])
	}

	disjoint := RsBfs(sr, sc)
	return len(stones) - disjoint

}

func RsBfs(sr, sc map[int][]int) int {
	count := 0
	for col, arr := range sc {
		count++
		delete(sc, col)
		for i := range arr {
			RsBfsRow(arr[i], sr, sc)
		}
	}
	return count
}

func RsBfsRow(row int, sr, sc map[int][]int) {
	arr := sr[row]
	delete(sr, row)
	for i := range arr {
		RsBfsCol(arr[i], sr, sc)
	}
}

func RsBfsCol(col int, sr, sc map[int][]int) {
	arr := sc[col]
	delete(sc, col)
	for i := range arr {
		RsBfsRow(arr[i], sr, sc)
	}
}
