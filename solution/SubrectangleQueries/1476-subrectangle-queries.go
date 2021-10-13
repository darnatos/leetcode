package SubrectangleQueries

type SubrectangleQueries struct {
	rect [][]int
	subs [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{rectangle, make([][]int, 0)}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	this.subs = append(this.subs, []int{row1, col1, row2, col2, newValue})
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	for i := len(this.subs) - 1; i >= 0; i-- {
		sub := this.subs[i]
		if row >= sub[0] && row <= sub[2] && col >= sub[1] && col <= sub[3] {
			return sub[4]
		}
	}
	return this.rect[row][col]
}
