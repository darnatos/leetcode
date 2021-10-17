package solution

func Candy2(ratings []int) int {
	res := len(ratings)
	up, down, peak := 0, 0, 0
	for i := 1; i < len(ratings); i++ {
		slope := ratings[i] - ratings[i-1]
		if slope > 0 {
			up++
			peak = up
			down = 0
			res += up
		} else if slope == 0 {
			up, down, peak = 0, 0, 0
		} else {
			up = 0
			down++
			res += down
			if down <= peak {
				res--
			}
		}
	}

	return res
}
