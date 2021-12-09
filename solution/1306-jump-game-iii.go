package solution

func CanReach(arr []int, start int) bool {
	if start < 0 || start >= len(arr) {
		return false
	}
	if arr[start] < 0 {
		return false
	}
	if arr[start] == 0 {
		return true
	}
	l, r := start-arr[start], start+arr[start]
	arr[start] = -1
	return CanReach(arr, l) || CanReach(arr, r)
}
