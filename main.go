package main

import (
	"fmt"
	"leetcode/solution"
)

func main() {
	//fmt.Println(solution.ReformatDate("6th Jun 1933"))
	//fmt.Println(solution.LongestSubsequence([]int{1, 5, 7, 8, 5, 3, 4, 2, 1}, -2))
	//fmt.Println(solution.LongestSubsequence([]int{1, 2, 3, 4}, 1))
	//fmt.Println(solution.RemoveDuplicates([]int{}))
	//nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	//k := solution.RemoveDuplicates(nums)
	//fmt.Println(k)
	//fmt.Println(nums[:k])
	//fmt.Println(solution.SmallestChair([][]int{{3, 10}, {1, 5}, {2, 6}}, 0), " should be ", 2)
	//fmt.Println(solution.SmallestChair([][]int{{33889, 98676}, {80071, 89737}, {44118, 52565}, {52992, 84310}, {78492, 88209}, {21695, 67063}, {84622, 95452}, {98048, 98856}, {98411, 99433}, {55333, 56548}, {65375, 88566}, {55011, 62821}, {48548, 48656}, {87396, 94825}, {55273, 81868}, {75629, 91467}}, 6), " should be ", 2)

	boxes := [...]int{2, 5, 10, 9, 4, 8, 6, 9, 9, 1}
	fmt.Println("expected", 16, "and got", solution.RemoveBoxes(boxes[:]))
}
