package solution

import "sort"

func IsBoomerang(points [][]int) bool {
	if points[0][0] == points[1][0] && points[1][0] == points[2][0] {
		return false
	} else if points[0][0] == points[1][0] {
		if points[0][1] == points[1][1] {
			return false
		}
		return true
	} else if points[1][0] == points[2][0] {
		if points[1][1] == points[2][1] {
			return false
		}
		return true
	} else if points[0][0] == points[2][0] {
		if points[0][1] == points[2][1] {
			return false
		}
		return true
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0] || points[i][0] == points[j][0] && points[i][1] < points[j][1]
	})
	return float64(points[1][1]-points[0][1])/float64(points[1][0]-points[0][0]) != float64(points[2][1]-points[1][1])/float64(points[2][0]-points[1][0])
}
