package solution

var adjList map[int][]int

func CanFinish(numCourses int, prerequisites [][]int) bool {
	adjList = make(map[int][]int, 16)
	for i := range prerequisites {
		adjList[prerequisites[i][1]] = append(adjList[prerequisites[i][1]], prerequisites[i][0])
	}

	courses := make([]int, numCourses)

	for i := range courses {
		if dfsAndReturnCycle(i, courses) {
			return false
		}
	}

	return true
}

func dfsAndReturnCycle(i int, courses []int) bool {
	if courses[i] == 1 {
		return true
	}
	if courses[i] == 2 {
		return false
	}
	courses[i] = 1

	for _, v := range adjList[i] {
		if dfsAndReturnCycle(v, courses) {
			return true
		}
	}
	courses[i] = 2

	return false
}
