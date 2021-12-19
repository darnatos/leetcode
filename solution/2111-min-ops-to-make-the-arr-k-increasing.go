package solution

func kIncreasing(arr []int, k int) int {
	res := 0
	for i := 0; i < k; i++ {
		subSeq := make([]int, 0)
		for j := i; j < len(arr); j += k {
			l, r := 0, len(subSeq)
			for l < r {
				m := l + (r-l)/2
				if subSeq[m] <= arr[j] {
					l = m + 1
				} else {
					r = m
				}
			}

			if l == len(subSeq) {
				subSeq = append(subSeq, arr[j])
			} else {
				subSeq[l] = arr[j]
			}
		}
		res += len(subSeq)
	}
	return len(arr) - res
}
