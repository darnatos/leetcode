package solution

func getWinner(arr []int, k int) int {
	if k > len(arr) {
		mx := 0
		for i := range arr {
			if mx < arr[i] {
				mx = arr[i]
			}
		}
		return mx
	}
	cnt := make(map[int]int)
	i := 0
	for j := 1; j < len(arr); {
		if arr[i] < arr[j] {
			cnt[arr[j]]++
			if cnt[arr[j]] == k {
				return arr[j]
			}
			i = j
		} else {
			cnt[arr[i]]++
			if cnt[arr[i]] == k {
				return arr[i]
			}
		}
		j++
	}
	return arr[i]
}
