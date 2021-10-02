package solution

func GetMinSwaps(num string, k int) int {
	seq := []byte(num)
	for i := 0; i < k; i++ {
		nextP(seq)
	}
	seq0 := []byte(num)
	i, j, res := 0, 0, 0
	for string(seq0) != string(seq) {
		for i < len(seq) && seq0[i] == seq[i] {
			i++
		}
		j = i + 1
		for j < len(seq) && seq[j] != seq0[i] {
			j++
		}
		for j > i {
			seq[j-1], seq[j] = seq[j], seq[j-1]
			res++
			j--
		}
	}

	return res
}

func nextP(seq []byte) {
	reverse := func(i, j int) {
		for ; i < j; i, j = i+1, j-1 {
			seq[i], seq[j] = seq[j], seq[i]
		}
	}
	k := len(seq) - 2
	for k >= 0 {
		if seq[k] < seq[k+1] {
			break
		}
		k--
	}
	if k == -1 {
		reverse(0, len(seq)-1)
	} else {
		for j := len(seq) - 1; j > k; j-- {
			if seq[k] < seq[j] {
				seq[k], seq[j] = seq[j], seq[k]
				break
			}
		}
		reverse(k+1, len(seq)-1)
	}
}
