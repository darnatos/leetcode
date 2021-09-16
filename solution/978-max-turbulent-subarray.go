package solution

func MaxTurbulenceSize(arr []int) int {
	if len(arr) <= 1 {
		return len(arr)
	}

	asc, l, r, res := 0, 0, 0, 0

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			r = i + 1
			if asc == 1 {
				l = i
				continue
			}
			asc = 1
			res = max(res, r-l)
		} else if arr[i] > arr[i+1] {
			r = i + 1
			if asc == -1 {
				l = i
				continue
			}
			asc = -1
			res = max(res, r-l)
		} else {
			l = i + 1
			asc = 0
		}
	}

	return res + 1
}
