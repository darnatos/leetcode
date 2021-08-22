package solution

func MinEatingSpeed(piles []int, h int) int {
	high := 0
	for i := range piles {
		if piles[i] > high {
			high = piles[i]
		}
	}

	low, mid, sum := 1, 0, 0
	for low < high {
		mid = low + ((high - low) >> 1)
		sum = 0
		for _, v := range piles {
			sum += 1 + (v-1)/mid
			if sum > h {
				goto invalid
			}
		}
		high = mid
		continue
	invalid:
		low = mid + 1
		continue
	}

	return low
}
