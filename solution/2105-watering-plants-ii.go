package solution

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	res := 0
	ac := capacityA
	bc := capacityB
	a, b := 0, len(plants)-1
	for a < b {
		if plants[a] > ac {
			res++
			ac = capacityA
		}
		ac -= plants[a]
		if plants[b] > bc {
			res++
			bc = capacityB
		}
		bc -= plants[b]
		a++
		b--
	}
	if a == b {
		if ac < plants[a] && bc < plants[b] {
			res++
		}
	}
	return res
}
