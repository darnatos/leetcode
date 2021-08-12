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
	fmt.Println(solution.SmallestChair([][]int{{3, 10}, {1, 5}, {2, 6}}, 0), " should be ", 2)
}
