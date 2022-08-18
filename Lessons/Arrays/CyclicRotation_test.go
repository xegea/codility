package arrays

import "testing"

type CyclicRotationTest struct {
	arg1     []int
	arg2     int
	expected []int
}

var cyclicRotationTests = []CyclicRotationTest{
	{[]int{8}, 10, []int{8}},
	{[]int{0, 0, 0}, 1, []int{0, 0, 0}},
	{[]int{1, 2, 3, 4}, 4, []int{1, 2, 3, 4}},
	{[]int{3, 8, 9, 7, 6}, 3, []int{9, 7, 6, 3, 8}},
}

func TestSolution(t *testing.T) {
	for _, test := range cyclicRotationTests {
		if output := Solution(test.arg1, test.arg2); !Equal(output, test.expected) {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}

func Equal(a, b []int) bool {
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
