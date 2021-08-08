package solution

import (
	"testing"
)

type input struct {
	nums     []int
	target   int
	expected []int
}

var testcases = []input{
	{[]int{3, 3}, 6, []int{0, 1}},
	{[]int{3, 2, 4}, 6, []int{1, 2}},
	{[]int{9, 8, 1, 2, 6, 4, 12}, 16, []int{5, 6}},
}

func TestTwoSum(t *testing.T) {
	for _, v := range testcases {
		expected := v.expected
		actual := TwoSum(v.nums, v.target)
		if !equal(expected, actual) {
			t.Errorf("Expected: %d, got %d", expected, actual)
		}
	}
}

func equal(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
