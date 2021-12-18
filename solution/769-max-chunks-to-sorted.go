package solution

func maxChunksToSorted1(arr []int) int {
	res := 0
	j, anchor := 0, 0
	for i := range arr {
		if j < arr[i] {
			j = arr[i]
		}
		if i == j {
			res++
			j = anchor + 1
		}
	}

	return res
}
