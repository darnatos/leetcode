package solution

type DetectSquares struct {
	points [][]int
}

func NewDetectSquares() DetectSquares {
	tmp := make([]int, 1001*1001)
	points := make([][]int, 1001)
	for i := range points {
		points[i] = tmp[i*1001 : i*1001+1001]
	}
	return DetectSquares{points}
}

func (dts *DetectSquares) Add(point []int) {
	dts.points[point[0]][point[1]]++
}

func (dts *DetectSquares) Count(point []int) int {
	res := 0
	x0, y0 := point[0], point[1]
	dx := [...]int{1, -1, 1, -1}
	dy := [...]int{1, 1, -1, -1}

	for z := 0; z < 4; z++ {
		x, y := x0, y0
		for true {
			x += dx[z]
			y += dy[z]
			if x < 0 || x > 1000 || y < 0 || y > 1000 {
				break
			}
			res += dts.points[x][y0] * dts.points[x][y] * dts.points[x0][y]
		}
	}

	return res
}
