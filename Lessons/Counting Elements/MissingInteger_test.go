package countingelements

import "testing"

type MissingIntegerTest struct {
	arg1     []int
	expected int
}

var missingIntegerTests = []MissingIntegerTest{
	{[]int{1, 2, 3, 4, 5}, 6},
	{[]int{1, 2, 4}, 3},
	{[]int{1, 2, 3}, 4},
	{[]int{-1, -3}, 1},
	{[]int{2}, 1},
}

func TestSolution(t *testing.T) {
	for _, test := range missingIntegerTests {
		if output := Solution(test.arg1); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
