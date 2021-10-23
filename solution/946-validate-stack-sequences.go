package solution

func ValidateStackSequences(pushed []int, popped []int) bool {
	i, j := 0, 0
	for _, x := range pushed {
		pushed[i] = x
		i++
		for i > 0 && popped[j] == pushed[i-1] {
			j++
			i--
		}
	}

	return i == 0
}
