package solution

import "strconv"

func DiffWaysToCompute(expr string) []int {
	if len(expr) < 3 {
		tmp, _ := strconv.Atoi(expr)
		return []int{tmp}
	}

	m := make([]int, 0, 2)

	for i := 0; i < len(expr); i++ {
		if expr[i] == '+' || expr[i] == '-' || expr[i] == '*' {
			p1 := DiffWaysToCompute(expr[:i])
			p2 := DiffWaysToCompute(expr[i+1:])
			for _, j := range p1 {
				for _, k := range p2 {
					switch expr[i] {
					case '+':
						m = append(m, j+k)
					case '-':
						m = append(m, j-k)
					case '*':
						m = append(m, j*k)
					}
				}
			}
		}
	}
	return m
}
