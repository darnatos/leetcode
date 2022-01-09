package solution

func isRobotBounded(instructions string) bool {
	x, y := 0, 0
	d := 0
	dir := []int{0, 1, 0, -1, 0}
	for i := range instructions {
		switch instructions[i] {
		case 'G':
			x += dir[d]
			y += dir[d+1]
		case 'L':
			d--
			if d == -1 {
				d = 3
			}
		case 'R':
			d++
			if d == 4 {
				d = 0
			}
		default:
			return false
		}
	}
	return x == 0 && y == 0 || d != 0
}
