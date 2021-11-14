package solution

type Robot struct {
	width  int
	height int
	dir    int
	dirs   []string
	x      int
	y      int
}

func ConstructRobot(width int, height int) Robot {
	return Robot{
		width:  width - 1,
		height: height - 1,
		dir:    0,
		dirs:   []string{"East", "North", "West", "South"},
		x:      0,
		y:      0,
	}
}

func (rbt *Robot) Move(num int) {
	num = (num-1)%((rbt.width+rbt.height)*2) + 1
	for num > 0 {
		dist := num
		if rbt.dir == 0 {
			if num >= rbt.width-rbt.x {
				dist = rbt.width - rbt.x
			}
			rbt.x += dist
		} else if rbt.dir == 1 {
			if num >= rbt.height-rbt.y {
				dist = rbt.height - rbt.y
			}
			rbt.y += dist
		} else if rbt.dir == 2 {
			if num >= rbt.x {
				dist = rbt.x
			}
			rbt.x -= dist
		} else {
			if num >= rbt.y {
				dist = rbt.y
			}
			rbt.y -= dist
		}
		num -= dist
		if num > 0 {
			rbt.dir = (rbt.dir + 1) % 4
		}
	}
}

func (rbt *Robot) GetPos() []int {
	return []int{rbt.x, rbt.y}
}

func (rbt *Robot) GetDir() string {
	return rbt.dirs[rbt.dir]
}

/**
 * Your Robot object will be instantiated and called as such:
 * obj := Constructor(width, height);
 * obj.Move(num);
 * param_2 := obj.GetPos();
 * param_3 := obj.GetDir();
 */

/**
 * Your Robot object will be instantiated and called as such:
 * obj := Constructor(width, height);
 * obj.Move(num);
 * param_2 := obj.GetPos();
 * param_3 := obj.GetDir();
 */
