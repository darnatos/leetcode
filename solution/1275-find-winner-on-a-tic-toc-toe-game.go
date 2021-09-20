package solution

import "fmt"

func Tictactoe(moves [][]int) string {
	graph := make([][]int, 3)
	for i := range graph {
		graph[i] = make([]int, 3)
	}
	for i, move := range moves {
		graph[move[0]][move[1]] = 1+(i&1)
	}
	fmt.Println(graph)

	for i := range graph {
		flag := graph[i][0]
		line := flag
		for j := 1; j < 3; j++ {
			if graph[i][j] != flag {
				line = 0
				break
			}
		}
		if line > 0 {
			if line == 1 {
				return "A"
			} else {
				return "B"
			}
		}
	}
	for j := range graph {
		flag := graph[0][j]
		line := flag
		for i := 1; i < 3; i++ {
			if graph[i][j] != flag {
				line = 0
				break
			}
		}
		if line > 0 {
			if line == 1 {
				return "A"
			} else {
				return "B"
			}
		}
	}
	if flag := graph[1][1]; flag > 0 {
		if graph[0][0] == flag && graph[2][2] == flag || graph[2][0] == flag && graph[0][2] == flag {
			if flag == 1 {
				return "A"
			} else {
				return "B"
			}
		}
	}
	if len(moves) == 9 {
		return "Draw"
	}
	return "Pending"
}

