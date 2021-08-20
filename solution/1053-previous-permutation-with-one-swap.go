package solution

func prevPermOpt1(arr []int) []int {
	x := -1
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i-1] > arr[i] {
			x = i - 1
			break
		}
	}
	if x == -1 {
		return arr
	}

	y := x + 1
	for j := y; j < len(arr); j++ {
		if arr[j] < arr[x] {
			if arr[j] > arr[y] {
				y = j
			}
		}
	}

	arr[x], arr[y] = arr[y], arr[x]

	return arr
}
