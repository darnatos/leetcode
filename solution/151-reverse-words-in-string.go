package solution

func ReverseWords(s string) string {
	arr := []byte(s)
	n := len(arr)
	i, j := 0, 0
	for j < n {
		for j < n && arr[j] == ' ' {
			j++
		}
		for j < n && arr[j] != ' ' {
			arr[i] = arr[j]
			i, j = i+1, j+1
		}
		for j < n && arr[j] == ' ' {
			j++
		}
		if j < n {
			arr[i] = ' '
			i++
		}
	}
	arr = arr[:i]

	for l, r := 0, len(arr)-1; l < r; l, r = l+1, r-1 {
		arr[l], arr[r] = arr[r], arr[l]
	}

	j = 0
	for i := 0; i <= len(arr); i++ {
		if i == len(arr) || arr[i] == ' ' {
			for l, r := j, i-1; l < r; l, r = l+1, r-1 {
				arr[l], arr[r] = arr[r], arr[l]
			}
			j = i + 1
		}
	}

	return string(arr)
}
