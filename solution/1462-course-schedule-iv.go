package solution

func CheckIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	tmp := make([]bool, numCourses*numCourses)
	connected := make([][]bool, numCourses)
	for i := range connected {
		connected[i] = tmp[i*numCourses : i*numCourses+numCourses]
	}

	for _, p := range prerequisites {
		connected[p[0]][p[1]] = true
	}

	for k := 0; k < numCourses; k++ {
		for i := 0; i < numCourses; i++ {
			for j := 0; j < numCourses; j++ {
				connected[i][j] = connected[i][j] || connected[i][k] && connected[k][j]
			}
		}
	}

	res := make([]bool, len(queries))
	for i, q := range queries {
		res[i] = connected[q[0]][q[1]]
	}

	return res
}
