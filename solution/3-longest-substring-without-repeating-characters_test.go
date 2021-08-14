package solution

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	var testcases = []struct {
		input  string
		output int
	}{{"abcabcbb", 3}, {"pwwkew", 3}, {"dvwd", 3}, {"dbqwewddbqwea", 6}}

	for _, testcase := range testcases {
		actual := LengthOfLongestSubstring(testcase.input)
		expected := testcase.output

		if expected != actual {
			t.Errorf("expected %d but found %d", expected, actual)
		}
	}
}
