package solution

func ReplaceElements(arr []int) []int {
	high := -1
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] > high {
			arr[i], high = high, arr[i]
		} else {
			arr[i] = high
		}
	}
	return arr
}
