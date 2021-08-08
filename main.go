package main

import (
	"fmt"
	"leetcode/solution"
)

func main() {
	//fmt.Println(solution.ReformatDate("6th Jun 1933"))
	fmt.Println(solution.LongestSubsequence([]int{1, 5, 7, 8, 5, 3, 4, 2, 1}, -2))
	fmt.Println(solution.LongestSubsequence([]int{1, 2, 3, 4}, 1))
}
