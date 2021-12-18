package solution

func maxChunksToSorted(arr []int) int {
	l := make([]int, len(arr))
	l[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		l[i] = max(l[i-1], arr[i])
	}
	r := arr[len(arr)-1]
	res := 1
	for i := len(arr) - 2; i >= 0; i-- {
		r = min(r, arr[i+1])
		if l[i] <= r {
			res++
		}
	}
	return res
}
