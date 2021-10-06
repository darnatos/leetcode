package solution

func LargestRectangleArea(heights []int) int {
	ret := 0
	r := 0
	stack := make([]int, 0, 2)
	for r < len(heights) {
		if len(stack) == 0 || heights[stack[len(stack)-1]] <= heights[r] {
			stack = append(stack, r)
			r++
		} else {
			tp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			w := r
			if len(stack) > 0 {
				w = r - stack[len(stack)-1] - 1
			}
			ret = max(ret, heights[tp]*w)
		}
	}
	for len(stack) > 0 {
		tp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		w := r
		if len(stack) > 0 {
			w = r - stack[len(stack)-1] - 1
		}
		ret = max(ret, heights[tp]*w)
	}

	return ret
}
