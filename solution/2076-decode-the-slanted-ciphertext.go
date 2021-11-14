package solution

func DecodeCiphertext(encodedText string, rows int) string {
	sb := make([]byte, 0)
	cols := (len(encodedText)-1)/rows + 1

	for i := 0; i < cols; i++ {
		for j := i; j < len(encodedText); j += cols + 1 {
			sb = append(sb, encodedText[j])
		}
	}
	end := 0
	for i := len(sb) - 1; i >= 0; i-- {
		if sb[i] != ' ' {
			end = i + 1
			break
		}
	}
	return string(sb[:end])
}
