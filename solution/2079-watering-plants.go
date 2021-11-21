package solution

func WateringPlants(plants []int, capacity int) int {
	cur := capacity
	res := len(plants)
	for i := 0; i < len(plants); i++ {
		if plants[i] > cur {
			res += 2 * i
			cur = capacity
		}
		cur -= plants[i]
	}
	return res
}
