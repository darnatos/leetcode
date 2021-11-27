package solution

func minCost(startPos []int, homePos []int, rowCosts []int, colCosts []int) int {
	dx, dy := 1, 1
	if homePos[0] < startPos[0] {
		dx = -1
	}
	if homePos[1] < startPos[1] {
		dy = -1
	}

	res := 0
	for i := startPos[0] + dx; i*dx <= homePos[0]*dx; i += dx {
		res += rowCosts[i]
	}
	for j := startPos[1] + dy; j*dy <= homePos[1]*dy; j += dy {
		res += colCosts[j]
	}

	return res
}
