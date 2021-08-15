package solution

import "math"

func LargestTriangleArea(points [][]int) float64 {
	var res float64
	l := len(points)
	for i := 0; i < l-2; i++ {
		for j := i + 1; j < l-1; j++ {
			for k := j + 1; k < l; k++ {
				res = math.Max(res, triangleArea(points[i], points[j], points[k]))
			}
		}
	}

	return res
}

func triangleArea(a, b, c []int) float64 {
	return math.Abs(float64((a[0]*b[1]+b[0]*c[1]+c[0]*a[1])-(a[1]*b[0]+b[1]*c[0]+c[1]*a[0])) / 2)
}
