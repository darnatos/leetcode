package main

import (
	"fmt"
	"leetcode/solution"
)

func main() {
	input := []int{3, 3}
	target := 6
	result := solution.TwoSum(input, target)
	fmt.Println(result)
}
